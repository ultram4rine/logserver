package service

import (
	"context"
	"errors"
	"net"
	"time"

	"git.sgu.ru/ultramarine/logserver/pb"

	"github.com/jmoiron/sqlx"
)

var facilityMap = map[uint8]string{
	0:  "kern",
	1:  "user",
	2:  "mail",
	3:  "daemon",
	4:  "auth",
	5:  "syslog",
	6:  "lpr",
	7:  "news",
	8:  "uucp",
	9:  "cron",
	10: "auth",
	11: "ftp",
	12: "ntp",
	13: "audit",
	14: "alert",
	15: "cron",
	16: "local0",
	17: "local1",
	18: "local2",
	19: "local3",
	20: "local4",
	21: "local5",
	22: "local6",
	23: "local7",
}

var severityMap = map[uint8]string{
	0: "Emergency",
	1: "Alert",
	2: "Critical",
	3: "Error",
	4: "Warning",
	5: "Notice",
	6: "Informational",
	7: "Debug",
}

// Service interface.
type Service interface {
	GetDHCPLogs(ctx context.Context, req *pb.DHCPLogsRequest) (*pb.DHCPLogsResponse, error)
	GetNginxLogs(ctx context.Context, req *pb.NginxLogsRequest) (*pb.NginxLogsResponse, error)
	GetNginxHosts(ctx context.Context, req *pb.NginxHostsRequest) (*pb.NginxHostsResponse, error)
	GetSwitchLogs(ctx context.Context, req *pb.SwitchLogsRequest) (*pb.SwitchLogsResponse, error)
	GetSimilarSwitches(ctx context.Context, req *pb.SimilarSwitchesRequest) (*pb.SimilarSwitchesResponse, error)
}

// LogService is a Service interface implementation.
type LogService struct {
	DB *sqlx.DB
}

// GetDHCPLogs returns DHCP logs from given MAC address and time interval.
func (s LogService) GetDHCPLogs(ctx context.Context, req *pb.DHCPLogsRequest) (*pb.DHCPLogsResponse, error) {
	timeFrom, timeTo := parseTime(req.From, req.To)

	rows, err := s.DB.QueryxContext(ctx, "SELECT ts, message, ip FROM dhcp.events WHERE mac = MACStringToNum(?) AND ts > ? AND ts < ? ORDER BY ts DESC", req.MAC, timeFrom, timeTo)
	if err != nil {
		return &pb.DHCPLogsResponse{}, err
	}

	var logs = new(pb.DHCPLogsResponse)
	for rows.Next() {
		var (
			l  = new(pb.DHCPLog)
			ts string
			ip net.IP
		)

		if err = rows.Scan(&ts, &l.Message, &ip); err != nil {
			return &pb.DHCPLogsResponse{}, err
		}

		t, err := time.Parse(time.RFC3339, ts)
		if err != nil {
			return &pb.DHCPLogsResponse{}, err
		}

		l.Timestamp = t.Format("02/01/2006 15:04:05")
		l.Ip = ip.String()

		logs.Logs = append(logs.Logs, l)
	}
	if rows.Err() != nil {
		return &pb.DHCPLogsResponse{}, err
	}

	return logs, nil
}

// GetNginxLogs return logs from given nginx host and time interval.
func (s LogService) GetNginxLogs(ctx context.Context, req *pb.NginxLogsRequest) (*pb.NginxLogsResponse, error) {
	timeFrom, timeTo := parseTime(req.From, req.To)

	const query = "SELECT timestamp, message, facility, severity FROM nginx WHERE host = ? AND timestamp > ? AND timestamp < ? ORDER BY timestamp DESC"
	rows, err := s.DB.QueryxContext(ctx, query, req.Hostname, timeFrom, timeTo)
	if err != nil {
		return &pb.NginxLogsResponse{}, err
	}

	var logs = new(pb.NginxLogsResponse)
	for rows.Next() {
		var (
			l        = new(pb.NginxLog)
			ts       string
			facility uint8
			severity uint8
		)

		if err = rows.Scan(&ts, &l.Message, &facility, &severity); err != nil {
			return &pb.NginxLogsResponse{}, err
		}

		t, err := time.Parse(time.RFC3339, ts)
		if err != nil {
			return &pb.NginxLogsResponse{}, err
		}

		l.Timestamp = t.Format("02/01/2006 15:04:05")

		logs.Logs = append(logs.Logs, l)
	}
	if rows.Err() != nil {
		return &pb.NginxLogsResponse{}, err
	}

	return logs, nil
}

// GetNginxHosts returns available for view logs nginx hosts.
func (s LogService) GetNginxHosts(ctx context.Context, req *pb.NginxHostsRequest) (*pb.NginxHostsResponse, error) {
	rows, err := s.DB.QueryxContext(ctx, "SELECT DISTINCT hostname FROM nginx")
	if err != nil {
		return &pb.NginxHostsResponse{}, err
	}

	var hosts = new(pb.NginxHostsResponse)
	for rows.Next() {
		var h = new(pb.NginxHost)

		if err = rows.Scan(&h.Name); err != nil {
			return &pb.NginxHostsResponse{}, err
		}

		hosts.Hosts = append(hosts.Hosts, h)
	}
	if rows.Err() != nil {
		return &pb.NginxHostsResponse{}, err
	}

	return hosts, nil
}

// GetSwitchLogs returns logs from given switch and time interval.
func (s LogService) GetSwitchLogs(ctx context.Context, req *pb.SwitchLogsRequest) (*pb.SwitchLogsResponse, error) {
	timeFrom, timeTo := parseTime(req.From, req.To)

	const query = "SELECT ts_local, ts_remote, log_msg, facility, severity FROM switchlogs WHERE sw_name = ? AND ts_local > ? AND ts_local < ? ORDER BY ts_local DESC"
	rows, err := s.DB.QueryxContext(ctx, query, req.Name, timeFrom, timeTo)
	if err != nil {
		return &pb.SwitchLogsResponse{}, err
	}

	var logs = new(pb.SwitchLogsResponse)
	for rows.Next() {
		var (
			l        = new(pb.SwitchLog)
			tsLocal  string
			tsRemote string
			facility uint8
			severity uint8
		)

		if err = rows.Scan(&tsLocal, &tsRemote, &l.Message, &facility, &severity); err != nil {
			return &pb.SwitchLogsResponse{}, err
		}

		tLocal, err := time.Parse(time.RFC3339, tsLocal)
		if err != nil {
			return &pb.SwitchLogsResponse{}, err
		}
		tRemote, err := time.Parse(time.RFC3339, tsRemote)
		if err != nil {
			return &pb.SwitchLogsResponse{}, err
		}

		l.TsLocal = tLocal.Format("02/01/2006 15:04:05")
		l.TsRemote = tRemote.Format("02/01/2006 15:04:05")

		var ok bool
		if l.Facility, ok = facilityMap[facility]; !ok {
			return &pb.SwitchLogsResponse{}, errors.New("No such facility code")
		}
		if l.Severity, ok = severityMap[severity]; !ok {
			return &pb.SwitchLogsResponse{}, errors.New("No such severity code")
		}

		logs.Logs = append(logs.Logs, l)
	}
	if rows.Err() != nil {
		return &pb.SwitchLogsResponse{}, err
	}

	return logs, nil
}

// GetSimilarSwitches returns available for view logs switches, which names are similar to given.
func (s LogService) GetSimilarSwitches(ctx context.Context, req *pb.SimilarSwitchesRequest) (*pb.SimilarSwitchesResponse, error) {
	rows, err := s.DB.QueryxContext(ctx, "SELECT DISTINCT sw_name, sw_ip FROM switchlogs WHERE sw_name LIKE ?", req.Name+"%")
	if err != nil {
		return &pb.SimilarSwitchesResponse{}, err
	}

	var switches = new(pb.SimilarSwitchesResponse)
	for rows.Next() {
		var (
			s  = new(pb.SimilarSwitch)
			IP net.IP
		)

		if err = rows.Scan(&s.Name, &IP); err != nil {
			return &pb.SimilarSwitchesResponse{}, err
		}

		s.IP = IP.String()

		switches.Switches = append(switches.Switches, s)
	}
	if rows.Err() != nil {
		return &pb.SimilarSwitchesResponse{}, err
	}

	return switches, nil
}

func parseTime(fromUnix, toUnix int64) (time.Time, time.Time) {
	from := time.Unix(fromUnix, 0)
	to := time.Unix(toUnix, 0)
	return from, to
}
