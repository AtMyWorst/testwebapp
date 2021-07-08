<template>
  <div class="flex flex-col mx-6">
    <div class="flex items-center">
      <div class="flex-grow">
        <h1 v-if="title" class="font-bold text-2xl">{{ title }}</h1>
        <h2 v-if="subtitle" class="font-medium">{{ subtitle }}</h2>
      </div>
    </div>
    <slot></slot>
  </div>
</template>

<script>
export default {
  props: {
    title: String,
    subtitle: String
  }
}
</script>
