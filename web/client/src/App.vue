<template>
    <v-app>
        <v-card flat>
            <v-toolbar color="primary" dark extended flat></v-toolbar>

            <v-card class="mx-auto" max-width="1300" style="margin-top: -64px;">
                <v-toolbar>
                    <v-select v-model="selection" :items="items" hide-details single-line
                              label="Select which logs you want to see"></v-select>

                    <v-spacer></v-spacer>

                    <v-text-field v-if="selection==='DHCP Logs'" hide-details single-line
                                  placeholder="XX:XX:XX:XX:XX:XX" label="MAC address of device"></v-text-field>

                    <v-autocomplete v-else-if="selection==='Switch Logs'" hide-details single-line label="Switch name"
                                    v-model="sw" :items="similarSwitches" :search-input.sync="search"
                                    :loading="isLoading"
                                    color="primary" hide-no-data placeholder="Start typing to Search"
                                    return-object></v-autocomplete>

                    <v-spacer></v-spacer>

                    <v-select v-model="time" :items="times" hide-details single-line
                              label="Select for which time you need logs"></v-select>

                    <v-spacer></v-spacer>

                    <v-btn color="primary" v-on:click="periodForm=true">Choose time period</v-btn>

                    <v-spacer></v-spacer>

                    <v-btn v-if="selection==='DHCP Logs'" color="primary" v-on:click="getDHCPLogs">Show DHCP logs
                    </v-btn>
                    <v-btn v-else-if="selection==='Switch Logs'" color="primary" v-on:click="getSwitchLogs">Show Switch
                        logs
                    </v-btn>
                </v-toolbar>
            </v-card>
        </v-card>

        <v-content>
            <v-container fluid>
                <v-data-table v-if="selection==='DHCP Logs'" fixed-header :headers="DHCPHeaders" sort-by="timestamp"
                              :items="DHCPLogs"></v-data-table>
                <v-data-table v-else-if="selection==='Switch Logs'" fixed-header :headers="switchHeaders"
                              sort-by="timestamp"
                              :items="switchLogs"></v-data-table>
            </v-container>
        </v-content>

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
                                        ref="menu"
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
                                        <v-btn text color="primary" @click="menuFromDate = false">Cancel</v-btn>
                                        <v-btn text color="primary" @click="$refs.menu.save(fromDate)">OK</v-btn>
                                    </v-date-picker>
                                </v-menu>
                            </v-col>

                            <v-col>
                                <v-menu ref="menu" v-model="menuFromTime" :close-on-content-click="false"
                                        :nudge-right="40"
                                        :return-value.sync="fromTime" transition="scale-transition" offset-y
                                        max-width="290px"
                                        min-width="290px">
                                    <template v-slot:activator="{ on }">
                                        <v-text-field v-model="fromTime" label="From this time"
                                                      readonly v-on="on"></v-text-field>
                                    </template>
                                    <v-time-picker v-if="menuFromTime" v-model="fromTime" full-width use-seconds
                                                   format="24hr"
                                                   @click:second="$refs.menu.save(fromTime)"></v-time-picker>
                                </v-menu>
                            </v-col>
                        </v-row>

                        <v-row>
                            <v-col>
                                <v-menu
                                        ref="menu"
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
                                        <v-btn text color="primary" @click="menuToDate = false">Cancel</v-btn>
                                        <v-btn text color="primary" @click="$refs.menu.save(toDate)">OK</v-btn>
                                    </v-date-picker>
                                </v-menu>
                            </v-col>

                            <v-col>
                                <v-menu ref="menu" v-model="menuToTime" :close-on-content-click="false"
                                        :nudge-right="40"
                                        :return-value.sync="toTime" transition="scale-transition" offset-y
                                        max-width="290px"
                                        min-width="290px">
                                    <template v-slot:activator="{ on }">
                                        <v-text-field v-model="toTime" label="To this time"
                                                      readonly v-on="on"></v-text-field>
                                    </template>
                                    <v-time-picker v-if="menuToTime" v-model="toTime" full-width use-seconds
                                                   format="24hr"
                                                   @click:second="$refs.menu.save(toTime)"></v-time-picker>
                                </v-menu>
                            </v-col>
                        </v-row>
                    </v-form>
                </v-card-text>
                <v-divider></v-divider>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="orange darken-1" @click="periodForm=false">Add</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </v-app>
</template>

<script>
    import {mdiCalendar, mdiClockOutline, mdiClose} from '@mdi/js';

    import {logServiceClient} from './log_grpc_web_pb';
    import {DHCPLogsRequest, SwitchLogsRequest, SimilarSwitchesRequest} from './log_pb';

    export default {
        data() {
            return {
                mdiCalendar: mdiCalendar,
                mdiClockOutline: mdiClockOutline,
                mdiClose: mdiClose,

                client: null,

                selection: 'DHCP Logs',
                items: ['DHCP Logs', 'Switch Logs'],
                mac: '',
                sw: '',

                fromDate: '',
                toDate: '',
                fromTime: '',
                toTime: '',

                time: 'Last 5 minutes',
                times: ['Last 5 minutes', 'Last 15 minutes', 'Last 30 minutes', 'Last 1 hour', 'Last 3 hours', 'Last 6 hours', 'Last 12 hours', 'Last 1 day', 'Last 3 days', 'Last week', 'Period'],
                period: false,
                periodForm: false,

                similarSwitches: [],
                search: null,
                isLoading: false,

                menuFromDate: false,
                menuToDate: false,
                menuFromTime: false,
                menuToTime: false,

                DHCPHeaders: [{text: "IP", align: "start", value: "ip"}, {
                    text: "Timestamp",
                    value: "timestamp"
                }, {text: "Message", value: "message"}],
                DHCPLogs: [],

                switchHeaders: [{text: "IP", align: "start", value: "ip"}, {text: "Name", value: "name"}, {
                    text: "Timestamp",
                    value: "timestamp"
                }, {text: "Message", value: "message"}],
                switchLogs: [],
            }
        },

        created: function () {
            this.client = new logServiceClient("http://localhost:8080", null, null);
        }
        ,

        methods: {
            getDHCPLogs: function () {
                let req = new DHCPLogsRequest();
                req.setMac(this.mac)
                req.setFrom(this.fromTime)
                req.setTo(this.toTime)
                this.client.getDHCPLogs(req, {}, (err, resp) => {
                    this.DHCPLogs = resp.toObject().logsList;
                });
            }
            ,

            getSwitchLogs: function () {
                let req = new SwitchLogsRequest()
                req.setName(this.sw)
                req.setFrom(this.fromTime)
                req.setTo(this.toTime)
                this.client.getSwitchLogs(req, {}, (err, resp) => {
                    this.switchLogs = resp.toObject().logsList;
                })
            }
            ,

            getSimilarSwitches: function () {
                let req = new SimilarSwitchesRequest()
                req.setName(this.sw)
                this.client.getSimilarSwitches(req, {}, (err, resp) => {
                    this.similarSwitches = resp.toObject().switchesList;
                })
            }
        }
        ,

        watch: {
            search(val) {
                // Items have already been loaded
                if (this.switches.length > 0) return

                if (this.isLoading) return

                this.isLoading = true

                // Lazily load input items
                fetch('https://api.publicapis.org/entries')
                    .then(res => res.json())
                    .then(res => {
                        const {count, entries} = res
                        this.count = count
                        this.entries = entries
                    })
                    .catch(err => {
                        console.log(err)
                    })
                    .finally(() => (this.isLoading = false))
            }
            ,
        }
        ,
    }
    ;
</script>
