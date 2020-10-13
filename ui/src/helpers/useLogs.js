import axios from "axios";
import { ref } from "@vue/composition-api";

import config from "@/config/config";

const DHCPLogsEndpoint = `${config.apiURL}/dhcp`;
const nginxLogsEndpoint = `${config.apiURL}/nginx`;
const switchLogsEndpoint = `${config.apiURL}/switches`;

export default function () {
  const DHCPLogs = ref([]);
  const NginxLogs = ref([]);
  const SwitchLogs = ref([]);

  const DHCPHeaders = [
    { text: "IP", align: "start", value: "ip" },
    {
      text: "Timestamp",
      value: "timestamp",
    },
    { text: "Message", value: "message" },
  ];
  const NginxHeaders = [
    {
      text: "Timestamp",
      value: "timestamp",
    },
    { text: "Message", value: "message" },
    { text: "Facility", value: "facility" },
    { text: "Severity", value: "severity" },
  ];
  const SwitchHeaders = [
    {
      text: "Local timestamp",
      value: "tsLocal",
    },
    {
      text: "Remote timestamp",
      value: "tsRemote",
    },
    { text: "Message", value: "message" },
    { text: "Facility", value: "facility" },
    { text: "Severity", value: "severity" },
  ];

  const getDHCPLogs = async (mac, from, to) => {
    mac = mac.toUpperCase();

    if (mac.length !== 17) {
      alert("Mac-address have incorrect length!");
      return;
    }

    try {
      const resp = await axios.post(DHCPLogsEndpoint, {
        MAC: mac,
        from: from,
        to: to,
      });
      return resp.data.logs;
    } catch (err) {
      console.log(err);
      return [];
    }
  };

  const getNginxLogs = async (hostname, from, to) => {
    try {
      const resp = await axios.post(nginxLogsEndpoint, {
        hostname: hostname,
        from: from,
        to: to,
      });
      return resp.data.logs;
    } catch (err) {
      console.log(err);
      return [];
    }
  };

  const getSwitchLogs = async (name, from, to) => {
    try {
      const resp = await axios.post(switchLogsEndpoint, {
        name: name,
        from: from,
        to: to,
      });
      return resp.data.logs;
    } catch (err) {
      console.log(err);
      return [];
    }
  };

  return {
    DHCPLogs,
    NginxLogs,
    SwitchLogs,

    DHCPHeaders,
    NginxHeaders,
    SwitchHeaders,

    getDHCPLogs,
    getNginxLogs,
    getSwitchLogs,
  };
}
