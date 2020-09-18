import axios from "axios";
import { ref } from "@vue/composition-api";

import config from "@/config/config";

const similarSwitchesEndpoint = `${config.apiURL}/similar`;

export default function () {
  const SimilarSwitches = ref([]);

  const getSimilarSwitches = async (name) => {
    try {
      const resp = await axios.post(similarSwitchesEndpoint, {
        name: name,
      });
      return resp.data.switches;
    } catch (err) {
      console.log(err);
      return [];
    }
  };

  return {
    SimilarSwitches,

    getSimilarSwitches,
  };
}
