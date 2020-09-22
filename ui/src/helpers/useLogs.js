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
      value: "ts",
    },
    { text: "Message", value: "message" },
  ];
  const SwitchHeaders = [
    { text: "IP", align: "start", value: "ip" },
    { text: "Name", value: "name" },
    {
      text: "Timestamp",
      value: "ts",
    },
    { text: "Message", value: "message" },
  ];

  const getDHCPLogs = async (mac, from, to) => {
    const macRegExp = /[0-9a-fA-F]{12}/;

    mac = mac.toLowerCase();
    mac = mac.replace(/\.|-|:/g, "");

    if (mac.length === 12) {
      if (!macRegExp.test(mac)) {
        alert("Wrong mac-address!");
        return;
      }
    } else {
      alert("Mac-address too long!");
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
