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
        <template #body-cell-link="props">
          <q-td class="bg-blue-1" :props="props">
            <q-btn
              v-if="props.value != ''"
              color="white"
              text-color="black"
              label="link"
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
import { useQuasar } from "quasar";

const gamesColumns = [
  {
    name: "name",
    label: "Name",
    field: "name",
    sortable: true,
    align: "left",
  },
  {
    name: "rating",
    label: "Rating",
    field: "rating",
    sortable: true,
    align: "left",
    format: (val, row) => {
      if (val === 0) {
        return "";
      }
      const floored = Math.floor(val);
      return `${floored}%`;
    },
  },
  {
    name: "link",
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
        //.get("http://localhost:9000/games/all")
        .get(process.env.API_URL)
        .then((response) => {
          games.value = response.data;
          //console.log(games.value);
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
