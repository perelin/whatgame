<template>
  <q-page class="flex flex-center">
    <img
      alt="Quasar logo"
      src="~assets/quasar-logo-vertical.svg"
      style="width: 200px; height: 200px"
    />
  </q-page>
</template>

<script>
import { defineComponent, ref } from "vue";
import { api } from "boot/axios";
import { useQuasar } from "quasar";

export default defineComponent({
  name: "PageIndex",
  setup() {
    const $q = useQuasar();
    const data = ref(null);

    function loadData() {
      api
        .get(
          "https://catalog.gamepass.com/sigls/v2?id=29a81209-df6f-41fd-a528-2ae6b91f719c&language=en-us&market=US"
        )
        .then((response) => {
          data.value = response.data;
          console.log(data.value);
        })
        .catch((err) => {
          console.log("kaputt", err);
        });
    }

    return { data, loadData };
  },
  mounted() {
    this.loadData();
  },
});
</script>
