<template>
    <div class="station-picker">
        <div class="dialog">
            <div class="close-button" v-on:click="onClose">
                <img alt="Close" src="../../assets/close.png" />
            </div>
        </div>
        <div class="header">
            <div class="title">Add Station</div>
            <div class="search">
                <input v-model="search" label="Search" @input="onSearch(search)" placeholder="Search Stations" />
            </div>
        </div>
        <div class="main">
            <div class="child" v-for="station in visible" v-bind:key="station.id">
                <StationPickerStation :station="station" :selected="selected === station.id" @selected="(ev) => onSelected(station)" />
            </div>
        </div>
        <PaginationControls :page="paging.number" :totalPages="totalPages()" @new-page="onNewPage" />
        <div class="footer">
            <div class="button" v-on:click="onAdd" v-bind:class="{ enabled: selected }">Add Station</div>
        </div>
    </div>
</template>

<script lang="ts">
import _ from "lodash";
import Vue, { PropType } from "vue";
import CommonComponents from "@/views/shared";
import StationPickerStation from "./StationPickerStation.vue";
import PaginationControls from "./PaginationControls.vue";
import { DisplayStation } from "@/store";

export default Vue.extend({
    name: "StationPicker",
    components: {
        ...CommonComponents,
        StationPickerStation,
        PaginationControls,
    },
    props: {
        stations: {
            type: Array as PropType<DisplayStation[]>,
            required: true,
        },
        filter: {
            type: Function as PropType<(station: DisplayStation) => boolean>,
            default: (station) => true,
        },
    },
    data(): {
        selected: any | null;
        paging: {
            number: number;
            size: number;
        };
        search: string;
    } {
        return {
            selected: null,
            paging: {
                number: 0,
                size: 6,
            },
            search: "",
        };
    },
    computed: {
        visible(): DisplayStation[] {
            const start = this.paging.number * this.paging.size;
            const end = start + this.paging.size;
            return this.filteredStations().slice(start, end);
        },
    },
    methods: {
        filteredStations(): DisplayStation[] {
            if (this.search.length === 0) {
                return this.stations.filter((station) => this.filter(station));
            }
            return _.filter(
                this.stations.filter((station) => this.filter(station)),
                (station) => {
                    return station.name.toLowerCase().indexOf(this.search.toLowerCase()) >= 0;
                }
            );
        },
        totalPages(): number {
            return Math.ceil(this.filteredStations().length / this.paging.size);
        },
        onSelected(station: DisplayStation): void {
            this.selected = station.id;
        },
        onClose(): void {
            this.$emit("close");
        },
        onAdd(): void {
            if (this.selected) {
                this.$emit("add", this.selected);
            }
        },
        onNewPage(page: number): void {
            this.paging.number = page;
        },
    },
});
</script>

<style scoped>
.station-picker {
    display: flex;
    flex-direction: column;
}
.station-picker .dialog {
    display: flex;
}
.dialog .close-button {
    margin-left: auto;
    cursor: pointer;
}
.station-picker .header {
    display: flex;
}
.header .title {
    font-weight: 500;
    font-size: 20px;
    color: #2c3e50;
}
.header .search {
    margin-left: auto;
}
.station-picker .main {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    min-height: 164px; /* Eh */
}
.main .child {
    margin-bottom: 10px;
}
.station-picker .pagination {
    display: flex;
}
.station-picker .footer {
    display: flex;
}
.footer .button {
    font-size: 18px;
    font-weight: bold;
    text-align: center;
    padding: 10px;
    margin: 20px 0 0 0px;
    border: 1px solid rgb(215, 220, 225);
    border-radius: 4px;
    color: #000000;
    background-color: #efefef;
}
.footer .button.enabled {
    color: #ffffff;
    background-color: #ce596b;
    cursor: pointer;
}
input {
    font-size: inherit;
    padding: 0.3em;
    border: 2px solid #efefef;
    border-radius: 3px;
}
.header {
    margin-bottom: 1em;
}
</style>
