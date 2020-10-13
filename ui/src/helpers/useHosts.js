import axios from "axios";
import { ref } from "@vue/composition-api";

import config from "@/config/config";

const nginxHostsEndpoint = `${config.apiURL}/hosts`;

export default function () {
  const NginxHosts = ref([]);

  const getNginxHosts = async () => {
    try {
      const resp = await axios.post(nginxHostsEndpoint, {});
      return resp.data.switches;
    } catch (err) {
      console.log(err);
      return [];
    }
  };

  return {
    NginxHosts,

    getNginxHosts,
  };
}
