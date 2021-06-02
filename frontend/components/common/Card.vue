
<template>
  <div
    class="flex-grow flex flex-col bg-white rounded-card shadow-md m-4"
    @click="$router.push(to)"
  >
    <div class="flex items-center">
      <div class="flex-grow py-4 px-6">
        <h1 v-if="title" class="font-bold text-2xl">{{ title }}</h1>
        <h2 v-if="subtitle" class="font-medium">{{ subtitle }}</h2>
      </div>
      <div class="p-6" v-if="arrow">
        <fa-icon :icon="['far', 'chevron-right']" />
      </div>
    </div>
    <div class="flex-grow flex items-center py-4 pt-0">
      <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    title: String,
    subtitle: String,
    arrow: Boolean,
    to: Object
  }
}
</script>