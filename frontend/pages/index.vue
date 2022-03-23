
<template>
  <div
    class="background h-screen text-white p-2 flex flex-col justify-center items-center"
  >
    <div class="flex flex-row items-center">
      <button class="my-2">
        <fa-icon size="5x" :icon="['far', 'globe-asia']" />
      </button>
      <h2 class="font-bold text-5xl pl-4">
        webml.app
      </h2>
    </div>
    <nuxt-link to="/admin"
      ><Button class="mt-5 w-56 bg-white hover:bg-white text-primary"
        >Get started!</Button
      ></nuxt-link
    >
  </div>
</template>

<style scoped>
.background {
  background: linear-gradient(45deg, #ff6900, #ff994f);
  padding-bottom: 6reml;
}
</style>

<script>
import Button from '~/components/common/Button.vue'

export default {
  components: {
    Button