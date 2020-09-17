import axios from "axios";
import { ref } from "@vue/composition-api";

import config from "@/config/config";

const similarSwitchesEndpoint = `${config.apiURL}/similar`;

export default function () {
  const similarSwitches = ref([]);

  const getSimilarSwitches = async (name) => {
    try {
      const resp = await axios.post(similarSwitchesEndpoint, {
        name: name,
      });
      return resp.data;
    } catch (err) {
      console.log(err);
      return [];
    }
  };

  return {
    similarSwitches,

    getSimilarSwitches,
  };
}
