package logserver

import (
	"context"
	"net"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

// Service interface.
type Service interface {
	GetDHCPLogs(ctx context.Context, mac, from, to string) (DHCPLogsResponse, error)
	GetSwitchLogs(ctx context.Context, name, from, to string) (SwitchLogsResponse, error)
	GetSimilarSwitches(ctx context.Context, name string) (SimilarSwitchesResponse, error)
}

// LogService is a Service interface implementation.
type LogService struct{ DB *sqlx.DB }

// GetDHCPLogs returns DHCP logs from given MAC address and time interval.
func (s LogService) GetDHCPLogs(ctx context.Context, mac, from, to string) (DHCPLogsResponse, error) {
	timeFrom, timeTo, err := parseTime(from, to)
	if err != nil {
		return DHCPLogsResponse{}, err
	}

	rows, err := s.DB.QueryxContext(ctx, "SELECT ts, message, ip FROM dhcp.events WHERE mac = MACStringToNum(?) AND ts > ? AND ts < ? ORDER BY ts DESC", mac, timeFrom, timeTo)
	if err != nil {
		return DHCPLogsResponse{}, err
	}

	var logs DHCPLogsResponse
	for rows.Next() {
		var l dhcpLog
		if err = rows.Scan(&l.TimeStamp, &l.Message, &l.IP); err != nil {
			return DHCPLogsResponse{}, err
		}

		logs.Logs = append(logs.Logs, l)
	}
	if rows.Err() != nil {
		return DHCPLogsResponse{}, err
	}

	return logs, nil
}

// GetSwitchLogs returns logs from given switch and time interval.
func (s LogService) GetSwitchLogs(ctx context.Context, name, from, to string) (SwitchLogsResponse, error) {
	timeFrom, timeTo, err := parseTime(from, to)
	if err != nil {
		return SwitchLogsResponse{}, err
	}

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

func parseTime(fromStr, toStr string) (time.Time, time.Time, error) {
	from, err := strconv.Atoi(fromStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	to, err := strconv.Atoi(toStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	fromDuration := time.Minute * -time.Duration(from)
	toDuration := time.Minute * -time.Duration(to)

	timeFrom := time.Now().Add(fromDuration)
	timeTo := time.Now().Add(toDuration)

	return timeFrom, timeTo, nil
}
