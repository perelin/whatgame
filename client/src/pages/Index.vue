<template>
  <q-page class="flex flex-center">
    <div class="q-pa-md">
      <q-table
        title="Game Pass Games"
        :rows="gamesRows"
        :columns="gamesColumns"
        row-key="id"
        :pagination="initialPagination"
        :filter="filter"
      >
        <template #body-cell-igdbLink="props">
          <q-td class="bg-blue-1" :props="props">
            <q-btn
              v-if="props.value != ''"
              color="white"
              text-color="black"
              label="IGDB"
              :href="props.value"
              target="_blank"
            />
          </q-td>
        </template>

        <template #body-cell-gpLink="props">
          <q-td class="bg-blue-1" :props="props">
            <q-btn
              v-if="props.value != ''"
              color="white"
              text-color="black"
              label="GamePass"
              :href="props.value"
              target="_blank"
            />
          </q-td>
        </template>

        <template v-slot:top-right>
          <q-input
            borderless
            dense
            debounce="300"
            v-model="filter"
            placeholder="Search"
          >
            <template v-slot:append>
              <q-icon name="search" />
            </template>
          </q-input>
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
    name: "name",
    label: "Name",
    field: "name",
    sortable: true,
    align: "left",
  },

  // {
  //   name: "gpreleasedate",
  //   label: "Release",
  //   field: "gpreleasedatetimestamp",
  //   sortable: true,
  //   align: "left",
  // },
  // {
  //   name: "igdbfirstreleasedatetimestamp",
  //   label: "Release",
  //   field: "igdbfirstreleasedatetimestamp",
  //   sortable: true,
  //   align: "left",
  // },
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
    name: "gpreleasedate",
    label: "Xbox Release",
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
    name: "rating",
    label: "Rating",
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
    name: "igdbLink",
    label: "IGDB",
    field: "igdburl",
    sortable: false,
  },
  {
    name: "gpLink",
    label: "GamePass",
    field: "gpid",
    sortable: false,
    format: (val, row) => {
      return `https://www.xbox.com/en-us/games/store/gamename/${val}`;
    },
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
          console.log(games.value);
        })
        .catch((err) => {
          console.log(err);
        });
    }

    return { games, loadData, filter: ref("") };
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
