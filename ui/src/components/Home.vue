<template>
  <div id="home">
    <v-card flat>
      <v-toolbar color="primary" dark extended flat></v-toolbar>

      <v-card class="mx-auto" max-width="1300" style="margin-top: -64px;">
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
            v-model="sw"
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
            v-on:click="periodForm=true"
          >Choose time period</v-btn>

          <v-spacer></v-spacer>

          <v-btn
            v-if="selection === 'DHCP Logs'"
            color="primary"
            v-on:click="getDHCPLogs"
          >Show DHCP logs</v-btn>
          <v-btn v-else-if="selection === 'Switch Logs'" color="primary" v-on:click="getSwitchLogs">
            Show Switch
            logs
          </v-btn>
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
          :headers="switchHeaders"
          sort-by="timestamp"
          :items="switchLogs"
        ></v-data-table>
      </v-container>
    </v-main>

    <v-dialog v-model="periodForm" max-width="500px">
      <v-card dark>
        <v-toolbar>
          <v-toolbar-title>Choose period of time for logs</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-btn icon @click="periodForm = false; fromTime = ''; toTime = ''">
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
                    <v-text-field v-model="fromDate" label="From this date" readonly v-on="on"></v-text-field>
                  </template>
                  <v-date-picker v-model="fromDate" no-title scrollable>
                    <v-spacer></v-spacer>
                    <v-btn text color="primary" @click="menuFromDate = false">Cancel</v-btn>
                    <v-btn text color="primary" @click="$refs.menuFD.save(fromDate)">OK</v-btn>
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
                    <v-text-field v-model="fromTime" label="From this time" readonly v-on="on"></v-text-field>
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
                    <v-text-field v-model="toDate" label="To this date" readonly v-on="on"></v-text-field>
                  </template>
                  <v-date-picker v-model="toDate" no-title scrollable>
                    <v-spacer></v-spacer>
                    <v-btn text color="primary" @click="menuToDate = false">Cancel</v-btn>
                    <v-btn text color="primary" @click="$refs.menuTD.save(toDate)">OK</v-btn>
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
                    <v-text-field v-model="toTime" label="To this time" readonly v-on="on"></v-text-field>
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
          <v-btn color="primary" @click="periodForm=false">Ok</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import { mdiCalendar, mdiClockOutline, mdiClose } from "@mdi/js";
import axios from "axios";

export default {
  name: "Home",

  data() {
    return {
      mdiCalendar: mdiCalendar,
      mdiClockOutline: mdiClockOutline,
      mdiClose: mdiClose,

      client: null,

      selection: "DHCP Logs",
      items: ["DHCP Logs", "Switch Logs"],
      mac: "",
      sw: "",

      fromDate: "",
      toDate: "",
      fromTime: "",
      toTime: "",

      time: "Last 5 minutes",
      times: [
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
        "Last week"
      ],
      period: false,
      periodForm: false,

      similarSwitches: [],
      search: null,
      isLoading: false,

      menuFromDate: false,
      menuToDate: false,
      menuFromTime: false,
      menuToTime: false,

      DHCPHeaders: [
        { text: "IP", align: "start", value: "ip" },
        {
          text: "Timestamp",
          value: "ts"
        },
        { text: "Message", value: "message" }
      ],
      DHCPLogs: [],

      switchHeaders: [
        { text: "IP", align: "start", value: "ip" },
        { text: "Name", value: "name" },
        {
          text: "Timestamp",
          value: "ts"
        },
        { text: "Message", value: "message" }
      ],
      switchLogs: []
    };
  },

  methods: {
    getDHCPLogs: function() {
      let dates = this.transformDates(),
        unixFrom = dates.unixFrom,
        unixTo = dates.unixTo;

      axios
        .post("/api/dhcp", {
          mac: this.mac,
          from: unixFrom,
          to: unixTo
        })
        .then(resp => {
          this.DHCPLogs = resp.data.logs;
        })
        .catch(err => {
          console.log(err);
        });
    },

    getSwitchLogs: function() {
      let dates = this.transformDates(),
        unixFrom = dates.unixFrom,
        unixTo = dates.unixTo;

      axios
        .post("/api/switch", {
          name: this.sw.name,
          from: unixFrom,
          to: unixTo
        })
        .then(resp => {
          this.SwitchLogs = resp.data.logs;
        })
        .catch(err => {
          console.log(err);
        });
    },

    transformDates: function() {
      let unixFrom, unixTo;

      if (this.time === "Period") {
        unixFrom =
          new Date(`${this.fromDate} ${this.fromTime}`).getTime() / 1000;
        unixTo = new Date(`${this.toDate} ${this.toTime}`).getTime() / 1000;
      } else {
        unixFrom = new Date();
        switch (this.time) {
          case "Last 5 minutes":
            unixFrom.setMinutes(unixFrom.getMinutes() - 5);
            break;
          case "Last 15 minutes":
            unixFrom.setMinutes(unixFrom.getMinutes() - 15);
            break;
          case "Last 30 minutes":
            unixFrom.setMinutes(unixFrom.getMinutes() - 30);
            break;
          case "Last hour":
            unixFrom.setHours(unixFrom.getHours() - 1);
            break;
          case "Last 3 hours":
            unixFrom.setHours(unixFrom.getHours() - 3);
            break;
          case "Last 6 hours":
            unixFrom.setHours(unixFrom.getHours() - 6);
            break;
          case "Last 12 hours":
            unixFrom.setHours(unixFrom.getHours() - 12);
            break;
          case "Last day":
            unixFrom.setHours(unixFrom.getHours() - 24);
            break;
          case "Last 3 days":
            unixFrom.setHours(unixFrom.getHours() - 3 * 24);
            break;
          case "Last week":
            unixFrom.setHours(unixFrom.getHours() - 7 * 24);
            break;
        }

        unixFrom = Math.round(unixFrom.getTime() / 1000);
        unixTo = Math.round(new Date().getTime() / 1000);
      }

      return { unixFrom, unixTo };
    }
  },

  computed: {
    similars() {
      return this.similarSwitches.map(sw => {
        const desc = `${sw.name} - ${sw.ip}`;
        return Object.assign({}, sw, { desc });
      });
    }
  },

  watch: {
    search(val) {
      if (this.similarSwitches.length > 0) return;

      if (this.isLoading) return;

      this.isLoading = true;

      axios
        .post("/api/similar", { name: this.sw })
        .then(resp => {
          this.similarSwitches = resp.data.switches;
        })
        .catch(err => {
          console.log(err);
        })
        .finally(() => (this.isLoading = false));
    }
  }
};
</script>
