import axios from "axios";
import { ref } from "@vue/composition-api";

import config from "@/config/config";

const DHCPLogsEndpoint = `${config.apiURL}/dhcp`;
const switchLogsEndpoint = `${config.apiURL}/switches`;

export default function () {
  const DHCPLogs = ref([]);
  const SwitchLogs = ref([]);

  const DHCPHeaders = [
    { text: "IP", align: "start", value: "ip" },
    {
      text: "Timestamp",
      value: "timestamp",
    },
    { text: "Message", value: "message" },
  ];
  const SwitchHeaders = [
    {
      text: "Local timestamp",
      value: "ts_local",
    },
    {
      text: "Remote timestamp",
      value: "ts_remote",
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
    SwitchLogs,

    DHCPHeaders,
    SwitchHeaders,

    getDHCPLogs,
    getSwitchLogs,
  };
}
