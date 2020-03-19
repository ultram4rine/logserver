package logserver

import (
	"context"
	"encoding/binary"
	"net"
	"strconv"
	"time"

	pb "git.sgu.ru/ultramarine/logserver/proto"
	"github.com/jmoiron/sqlx"
)

// Service interface.
type Service interface {
	GetDHCPLogs(ctx context.Context, mac uint64, from, to string) (pb.DHCPLogs, error)
	GetSwitchLogs(ctx context.Context, name, from, to string) (pb.SwitchLogs, error)
	GetSimilarSwitches(ctx context.Context, name string) (pb.Switches, error)
}

// LogService is a Service interface implementation.
type LogService struct{ DB *sqlx.DB }

// GetDHCPLogs returns DHCP logs from given MAC address and time interval.
func (s LogService) GetDHCPLogs(ctx context.Context, mac uint64, from, to string) (pb.DHCPLogs, error) {
	timeFrom, timeTo, err := parseTime(from, to)
	if err != nil {
		return nil, err
	}

	var mhex []byte
	binary.BigEndian.PutUint64(mhex, mac)
	mcvt := net.HardwareAddr(mhex).String()

	rows, err := s.DB.QueryxContext(ctx, "SELECT ts, message, ip FROM dhcp.events WHERE mac = MACStringToNum(?) AND ts > ? AND ts < ? ORDER BY ts DESC", mcvt, timeFrom, timeTo)
	if err != nil {
		return nil, err
	}

	var logs *pb.DHCPLogs
	for rows.Next() {
		var l *pb.DHCPLog
		if err = rows.Scan(&l.Timestamp, &l.Message, &l.Ip); err != nil {
			return nil, err
		}

		logs.Log = append(logs.Log, l)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return logs, nil
}

// GetSwitchLogs returns logs from given switch and time interval.
func (s LogService) GetSwitchLogs(ctx context.Context, name, from, to string) (pb.SwitchLogs, error) {
	timeFrom, timeTo, err := parseTime(from, to)
	if err != nil {
		return nil, err
	}

	rows, err := s.DB.QueryxContext(ctx, "SELECT ts_remote, log_msg FROM switchlogs WHERE sw_name = ? AND ts_local > ? AND ts_local < ? ORDER BY ts_local DESC", name, timeFrom, timeTo)
	if err != nil {
		return nil, err
	}

	var logs *pb.SwitchLogs
	for rows.Next() {
		var l *pb.SwitchLog
		if err = rows.Scan(&l.Ts, &l.Message); err != nil {
			return nil, err
		}

		logs.Log = append(logs.Log, l)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return logs, nil
}

// GetSimilarSwitches returns available for view logs switches, which names are similar to given.
func (s LogService) GetSimilarSwitches(ctx context.Context, name string) (pb.Switches, error) {
	rows, err := s.DB.QueryxContext(ctx, "SELECT DISTINCT sw_name, sw_ip FROM switchlogs WHERE sw_name LIKE ?", name+"%")
	if err != nil {
		return nil, err
	}

	var switches *pb.Switches
	for rows.Next() {
		var s *pb.Switch
		if err = rows.Scan(&s.Name, &s.IP); err != nil {
			return nil, err
		}

		switches.Switch = append(switches.Switch, s)
	}
	if rows.Err() != nil {
		return nil, err
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
