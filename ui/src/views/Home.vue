<template>
  <div id="home">
    <v-card flat>
      <v-toolbar color="primary" dark extended flat></v-toolbar>

      <v-card class="mx-auto" max-width="1300" style="margin-top: -64px">
        <v-toolbar>
          <v-select
            v-model="selection"
            :items="items"
            hide-details
            single-line
            label="Select which logs you want to see"
          ></v-select>

          <v-spacer></v-spacer>

          <v-text-field
            v-if="selection === 'DHCP Logs'"
            v-model="mac"
            hide-details
            single-line
            placeholder="XX:XX:XX:XX:XX:XX"
            label="MAC address of device"
          ></v-text-field>

          <v-autocomplete
            v-else-if="selection === 'Switch Logs'"
            hide-details
            single-line
            label="Switch name"
            v-model="name"
            :items="similars"
            :search-input.sync="search"
            :loading="isLoading"
            item-text="desc"
            item-value="name"
            color="primary"
            hide-no-data
            placeholder="Start typing to Search"
            return-object
          ></v-autocomplete>

          <v-spacer></v-spacer>

          <v-select
            v-model="time"
            :items="times"
            hide-details
            single-line
            label="Select for which time you need logs"
          ></v-select>

          <v-spacer></v-spacer>

          <v-btn
            v-if="time === 'Period'"
            color="primary"
            v-on:click="periodForm = true"
            >Choose time period</v-btn
          >

          <v-spacer></v-spacer>

          <v-btn
            v-if="selection === 'DHCP Logs'"
            color="primary"
            v-on:click="insertDHCPLogs"
            >Show logs</v-btn
          >
          <v-btn
            v-else-if="selection === 'Switch Logs'"
            color="primary"
            v-on:click="insertSwitchLogs"
            >Show logs</v-btn
          >

          <v-spacer></v-spacer>

          <v-menu bottom left offset-y>
            <template v-slot:activator="{ on, attrs }">
              <v-btn icon v-bind="attrs" v-on="on">
                <v-icon>{{ mdiDotsVertical }}</v-icon>
              </v-btn>
            </template>

            <v-list>
              <v-btn text small v-on:click="logout">Logout</v-btn>
            </v-list>
          </v-menu>
        </v-toolbar>
      </v-card>
    </v-card>

    <v-main>
      <v-container fluid>
        <v-data-table
          v-if="selection === 'DHCP Logs'"
          fixed-header
          :headers="DHCPHeaders"
          sort-by="timestamp"
          :items="DHCPLogs"
        ></v-data-table>
        <v-data-table
          v-else-if="selection === 'Switch Logs'"
          fixed-header
          :headers="SwitchHeaders"
          sort-by="timestamp"
          :items="SwitchLogs"
        ></v-data-table>
      </v-container>
    </v-main>

    <v-dialog v-model="periodForm" max-width="500px">
      <v-card dark>
        <v-toolbar>
          <v-toolbar-title>Choose period of time for logs</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-btn icon @click="periodForm = false">
            <v-icon>{{ mdiClose }}</v-icon>
          </v-btn>
        </v-toolbar>
        <v-card-text>
          <v-form ref="form">
            <v-row>
              <v-col>
                <v-menu
                  ref="menuFD"
                  v-model="menuFromDate"
                  :close-on-content-click="false"
                  :return-value.sync="fromDate"
                  transition="scale-transition"
                  offset-y
                  min-width="290px"
                >
                  <template v-slot:activator="{ on }">
                    <v-text-field
                      v-model="fromDate"
                      label="From this date"
                      readonly
                      v-on="on"
                    ></v-text-field>
                  </template>
                  <v-date-picker v-model="fromDate" no-title scrollable>
                    <v-spacer></v-spacer>
                    <v-btn text color="primary" @click="menuFromDate = false"
                      >Cancel</v-btn
                    >
                    <v-btn
                      text
                      color="primary"
                      @click="$refs.menuFD.save(fromDate)"
                      >OK</v-btn
                    >
                  </v-date-picker>
                </v-menu>
              </v-col>

              <v-col>
                <v-menu
                  ref="menuFT"
                  v-model="menuFromTime"
                  :close-on-content-click="false"
                  :nudge-right="40"
                  :return-value.sync="fromTime"
                  transition="scale-transition"
                  offset-y
                  max-width="290px"
                  min-width="290px"
                >
                  <template v-slot:activator="{ on }">
                    <v-text-field
                      v-model="fromTime"
                      label="From this time"
                      readonly
                      v-on="on"
                    ></v-text-field>
                  </template>
                  <v-time-picker
                    v-if="menuFromTime"
                    v-model="fromTime"
                    full-width
                    use-seconds
                    format="24hr"
                    @click:second="$refs.menuFT.save(fromTime)"
                  ></v-time-picker>
                </v-menu>
              </v-col>
            </v-row>

            <v-row>
              <v-col>
                <v-menu
                  ref="menuTD"
                  v-model="menuToDate"
                  :close-on-content-click="false"
                  :return-value.sync="toDate"
                  transition="scale-transition"
                  offset-y
                  min-width="290px"
                >
                  <template v-slot:activator="{ on }">
                    <v-text-field
                      v-model="toDate"
                      label="To this date"
                      readonly
                      v-on="on"
                    ></v-text-field>
                  </template>
                  <v-date-picker v-model="toDate" no-title scrollable>
                    <v-spacer></v-spacer>
                    <v-btn text color="primary" @click="menuToDate = false"
                      >Cancel</v-btn
                    >
                    <v-btn
                      text
                      color="primary"
                      @click="$refs.menuTD.save(toDate)"
                      >OK</v-btn
                    >
                  </v-date-picker>
                </v-menu>
              </v-col>

              <v-col>
                <v-menu
                  ref="menuTT"
                  v-model="menuToTime"
                  :close-on-content-click="false"
                  :nudge-right="40"
                  :return-value.sync="toTime"
                  transition="scale-transition"
                  offset-y
                  max-width="290px"
                  min-width="290px"
                >
                  <template v-slot:activator="{ on }">
                    <v-text-field
                      v-model="toTime"
                      label="To this time"
                      readonly
                      v-on="on"
                    ></v-text-field>
                  </template>
                  <v-time-picker
                    v-if="menuToTime"
                    v-model="toTime"
                    full-width
                    use-seconds
                    format="24hr"
                    @click:second="$refs.menuTT.save(toTime)"
                  ></v-time-picker>
                </v-menu>
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" @click="periodForm = false">Ok</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import { mdiDotsVertical, mdiClose } from "@mdi/js";
import { ref, computed, watch } from "@vue/composition-api";

import useLogs from "@/helpers/useLogs";
import useSwitches from "@/helpers/useSwitches";

export default {
  name: "Home",

  setup() {
    const {
      DHCPLogs,
      SwitchLogs,
      DHCPHeaders,
      SwitchHeaders,
      getDHCPLogs,
      getSwitchLogs,
    } = useLogs();
    const { SimilarSwitches, getSimilarSwitches } = useSwitches();

    const selection = ref("DHCP Logs");
    const items = ["DHCP Logs", "Switch Logs"];

    const mac = ref("");
    const name = ref("");

    const fromDate = ref("");
    const toDate = ref("");
    const fromTime = ref("");
    const toTime = ref("");

    const time = ref("Last 5 minutes");
    const times = [
      "Period",
      "Last 5 minutes",
      "Last 15 minutes",
      "Last 30 minutes",
      "Last hour",
      "Last 3 hours",
      "Last 6 hours",
      "Last 12 hours",
      "Last day",
      "Last 3 days",
      "Last week",
    ];
    const period = ref(false);
    const periodForm = ref(false);

    const search = ref(null);
    const isLoading = ref(false);

    const menuFromDate = ref(false);
    const menuToDate = ref(false);
    const menuFromTime = ref(false);
    const menuToTime = ref(false);

    const insertDHCPLogs = () => {
      let dates = transformDates();

      getDHCPLogs(mac.value, dates.unixFrom, dates.unixTo).then(
        (logs) => (DHCPLogs.value = logs)
      );
    };
    const insertSwitchLogs = () => {
      let dates = transformDates();

      getSwitchLogs(name.value.name, dates.unixFrom, dates.unixTo).then(
        (logs) => (SwitchLogs.value = logs)
      );
    };

    const transformDates = () => {
      let unixFrom, unixTo;

      if (time.value === times[0]) {
        unixFrom =
          new Date(`${fromDate.value} ${fromTime.value}`).getTime() / 1000;
        unixTo = new Date(`${toDate.value} ${toTime.value}`).getTime() / 1000;
      } else {
        unixFrom = new Date();
        switch (time.value) {
          case times[1]:
            unixFrom.setMinutes(unixFrom.getMinutes() - 5);
            break;
          case times[2]:
            unixFrom.setMinutes(unixFrom.getMinutes() - 15);
            break;
          case times[3]:
            unixFrom.setMinutes(unixFrom.getMinutes() - 30);
            break;
          case times[4]:
            unixFrom.setHours(unixFrom.getHours() - 1);
            break;
          case times[5]:
            unixFrom.setHours(unixFrom.getHours() - 3);
            break;
          case times[6]:
            unixFrom.setHours(unixFrom.getHours() - 6);
            break;
          case times[7]:
            unixFrom.setHours(unixFrom.getHours() - 12);
            break;
          case times[8]:
            unixFrom.setHours(unixFrom.getHours() - 24);
            break;
          case times[9]:
            unixFrom.setHours(unixFrom.getHours() - 3 * 24);
            break;
          case times[10]:
            unixFrom.setHours(unixFrom.getHours() - 7 * 24);
            break;
        }

        unixFrom = Math.round(unixFrom.getTime() / 1000);
        unixTo = Math.round(new Date().getTime() / 1000);
      }

      return { unixFrom, unixTo };
    };

    const similars = computed(() => {
      return SimilarSwitches.value.map((sw) => {
        const desc = `${sw.name} - ${sw.IP}`;
        return Object.assign({}, sw, { desc });
      });
    });

    watch(
      () => search.value,
      () => {
        if (SimilarSwitches.value.length > 0) return;

        if (isLoading.value) return;

        isLoading.value = true;

        getSimilarSwitches(name.value).then((switches) => {
          SimilarSwitches.value = switches;
          isLoading.value = false;
        });
      }
    );

    return {
      DHCPLogs,
      SwitchLogs,
      SimilarSwitches,

      DHCPHeaders,
      SwitchHeaders,

      getSimilarSwitches,

      mac,
      name,
      fromDate,
      toDate,
      fromTime,
      toTime,

      selection,
      items,
      time,
      times,
      period,
      periodForm,

      search,
      isLoading,

      menuFromDate,
      menuToDate,
      menuFromTime,
      menuToTime,

      insertDHCPLogs,
      insertSwitchLogs,

      similars,

      mdiDotsVertical,
      mdiClose,
    };
  },

  methods: {
    logout: function () {
      this.$store.dispatch("AUTH_LOGOUT").then(() => {
        this.$router.push("/login");
      });
    },
  },
};
</script>
