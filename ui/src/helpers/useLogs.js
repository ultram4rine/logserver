import axios from "axios";
import { ref } from "@vue/composition-api";

import config from "@/config";

const DHCPLogsEndpoint = `${config.apiURL}/dhcp`;
const switchLogsEndpoint = `${config.apiURL}/switches`;

export default function () {
  const DHCPLogs = ref([]);
  const switchLogs = ref([]);

  const getDHCPLogs = async (mac, from, to) => {
    try {
      const resp = await axios.post(DHCPLogsEndpoint, {
        MAC: mac,
        from: from,
        to: to,
      });
      return resp.data;
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
      return resp.data;
    } catch (err) {
      console.log(err);
      return [];
    }
  };

  return {
    DHCPLogs,
    switchLogs,

    getDHCPLogs,
    getSwitchLogs,
  };
}
