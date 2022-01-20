<template>
  <q-page>
    <div class="q-pa-md">
      <q-banner rounded>
        What Game Pass game to play tonight? A simple overview of all currently
        available GamePass games with their score and release date. If you like
        to give feedback or have good ideas how to improve this, feel free to
        drop a note!
        <template v-slot:action>
          <q-btn
            href="https://forms.office.com/Pages/ResponsePage.aspx?id=DQSIkWdsW0yxEjajBLZtrQAAAAAAAAAAAANAAYBrVNdUQzNYNUVWMUFYTk4zUTVOSEkzQVdHQUU4Ri4u"
            target="_blank"
            label="Feedback and Requests"
            color="primary"
            class="text-black"
          />
        </template>
      </q-banner>
    </div>
    <div class="q-pa-md">
      <q-table
        :dense="$q.screen.lt.md"
        title="Game Pass Games"
        :rows="gamesRows"
        :columns="gamesColumns"
        row-key="id"
        :pagination="initialPagination"
        :filter="filter"
        :visible-columns="visibleColumns"
      >
        <!-- button to IGDB -->
        <template #body-cell-igdbLink="props">
          <q-td :props="props" auto-width>
            <q-btn
              v-if="props.value != ''"
              text-color="black"
              label="IGDB"
              :href="props.value"
              target="_blank"
              color="primary"
              class="text-black"
            />
          </q-td>
        </template>

        <!-- button to MSGP -->
        <template #body-cell-gpLink="props">
          <q-td :props="props">
            <q-btn
              v-if="props.value != ''"
              text-color="black"
              :href="props.value"
              target="_blank"
              color="primary"
              class="text-black"
              icon="link"
            />
          </q-td>
        </template>

        <!-- search and columns -->
        <template v-slot:top-right>
          <!-- search -->
          <q-input
            dense
            debounce="300"
            v-model="filter"
            placeholder="Search"
            outlined
          >
            <template v-slot:append>
              <q-icon name="search" />
            </template>
          </q-input>

          <!-- columns -->
          <q-select
            v-model="visibleColumns"
            multiple
            outlined
            dense
            options-dense
            :display-value="$q.lang.table.columns"
            emit-value
            map-options
            :options="gamesColumns"
            option-value="name"
            options-cover
            color="secondary"
          >
            <template
              v-slot:option="{ itemProps, opt, selected, toggleOption }"
            >
              <q-item v-bind="itemProps">
                <q-item-section>
                  <q-item-label v-html="opt.label" />
                </q-item-section>
                <q-item-section side>
                  <q-toggle
                    :model-value="selected"
                    color="secondary"
                    @update:model-value="toggleOption(opt)"
                  />
                </q-item-section>
              </q-item>
            </template>
          </q-select>
        </template>
      </q-table>
    </div>
  </q-page>
</template>

<script>
import { defineComponent, ref } from "vue";
import { api } from "boot/axios";
import { useQuasar, date } from "quasar";

const gamesColumns = [
  {
    name: "gpLink",
    label: "",
    field: "gpid",
    sortable: false,
    format: (val, row) => {
      return `https://www.xbox.com/en-us/games/store/gamename/${val}`;
    },
  },
  {
    name: "gpRating",
    label: "Rating",
    field: "gpaverageratingalltime",
    sortable: true,
    align: "left",
    format: (val, row) => {
      if (val === 0) {
        return "NA";
      }
      // const floored = Math.floor(val);
      // return `${floored}%`;
      return val;
    },
  },
  {
    name: "gpreleasedate",
    label: "Year",
    field: "gpreleasedatetimestamp",
    sortable: true,
    align: "left",
    format: (val, row) => {
      if (val === 0) {
        return "NA";
      }
      const releaseDate = new Date(val * 1000);
      return date.formatDate(releaseDate, "YYYY-MM");
    },
  },
  {
    name: "name",
    label: "Name",
    field: "name",
    sortable: true,
    align: "left",
    required: true,
  },
  {
    name: "rating",
    label: "IGDB Rating",
    field: "rating",
    sortable: true,
    align: "left",
    format: (val, row) => {
      if (val === 0) {
        return "NA";
      }
      const floored = Math.floor(val);
      return `${floored}%`;
    },
  },

  {
    name: "igdbfirstreleasedate",
    label: "Original Release",
    field: "igdbfirstreleasedatetimestamp",
    sortable: true,
    align: "left",
    format: (val, row) => {
      if (val === 0) {
        return "NA";
      }
      const releaseDate = new Date(val * 1000);
      return date.formatDate(releaseDate, "YYYY-MM");
    },
  },

  {
    name: "igdbLink",
    label: "IGDB",
    field: "igdburl",
    sortable: false,
  },
];
const gamesRows = [];

export default defineComponent({
  name: "PageIndex",
  data() {
    return {
      gamesColumns,
      gamesRows,
      initialPagination: {
        rowsPerPage: 0,
      },
    };
  },
  setup() {
    const $q = useQuasar();
    const games = ref(null);

    function loadData() {
      api
        .get(process.env.API_URL)
        .then((response) => {
          games.value = response.data;
          //console.log(games.value);
        })
        .catch((err) => {
          console.log(err);
        });
    }

    var visibleColumns = ref(["gpLink", "name", "gpRating", "gpreleasedate"]);

    return { games, loadData, filter: ref(""), visibleColumns };
  },
  watch: {
    games(newGames, oldGames) {
      console.log(newGames);
      this.gamesRows = newGames;
    },
  },
  mounted() {
    this.loadData();
  },
});
</script>
