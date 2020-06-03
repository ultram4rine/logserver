package logserver

import (
	"context"
	"net"
	"time"

	"github.com/jmoiron/sqlx"
)

// Service interface.
type Service interface {
	GetDHCPLogs(ctx context.Context, mac string, from, to int64) (DHCPLogsResponse, error)
	GetSwitchLogs(ctx context.Context, name string, from, to int64) (SwitchLogsResponse, error)
	GetSimilarSwitches(ctx context.Context, name string) (SimilarSwitchesResponse, error)
}

// LogService is a Service interface implementation.
type LogService struct {
	DB *sqlx.DB
}

// GetDHCPLogs returns DHCP logs from given MAC address and time interval.
func (s LogService) GetDHCPLogs(ctx context.Context, mac string, from, to int64) (DHCPLogsResponse, error) {
	timeFrom, timeTo := parseTime(from, to)

	rows, err := s.DB.QueryxContext(ctx, "SELECT ts, message, ip FROM dhcp.events WHERE mac = MACStringToNum(?) AND ts > ? AND ts < ? ORDER BY ts DESC", mac, timeFrom, timeTo)
	if err != nil {
		return DHCPLogsResponse{}, err
	}

	var logs DHCPLogsResponse
	for rows.Next() {
		var (
			l  dhcpLog
			ip net.IP
		)
		if err = rows.Scan(&l.TimeStamp, &l.Message, &ip); err != nil {
			return DHCPLogsResponse{}, err
		}

		l.IP = ip.String()

		logs.Logs = append(logs.Logs, l)
	}
	if rows.Err() != nil {
		return DHCPLogsResponse{}, err
	}

	return logs, nil
}

// GetSwitchLogs returns logs from given switch and time interval.
func (s LogService) GetSwitchLogs(ctx context.Context, name string, from, to int64) (SwitchLogsResponse, error) {
	timeFrom, timeTo := parseTime(from, to)

	rows, err := s.DB.QueryxContext(ctx, "SELECT ts_remote, log_msg FROM switchlogs WHERE sw_name = ? AND ts_local > ? AND ts_local < ? ORDER BY ts_local DESC", name, timeFrom, timeTo)
	if err != nil {
		return SwitchLogsResponse{}, err
	}

	var logs SwitchLogsResponse
	for rows.Next() {
		var l switchLog
		if err = rows.Scan(&l.TimeStamp, &l.Message); err != nil {
			return SwitchLogsResponse{}, err
		}

		logs.Logs = append(logs.Logs, l)
	}
	if rows.Err() != nil {
		return SwitchLogsResponse{}, err
	}

	return logs, nil
}

// GetSimilarSwitches returns available for view logs switches, which names are similar to given.
func (s LogService) GetSimilarSwitches(ctx context.Context, name string) (SimilarSwitchesResponse, error) {
	rows, err := s.DB.QueryxContext(ctx, "SELECT DISTINCT sw_name, sw_ip FROM switchlogs WHERE sw_name LIKE ?", name+"%")
	if err != nil {
		return SimilarSwitchesResponse{}, err
	}

	var switches SimilarSwitchesResponse
	for rows.Next() {
		var (
			s  similarSwitch
			IP net.IP
		)

		if err = rows.Scan(&s.Name, &IP); err != nil {
			return SimilarSwitchesResponse{}, err
		}

		s.IP = IP.String()

		switches.Sws = append(switches.Sws, s)
	}
	if rows.Err() != nil {
		return SimilarSwitchesResponse{}, err
	}

	return switches, nil
}

func parseTime(fromUnix, toUnix int64) (time.Time, time.Time) {
	from := time.Unix(fromUnix, 0)
	to := time.Unix(toUnix, 0)
	return from, to
}
