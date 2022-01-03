
<template>
  <div>
    <Header :title="title" subtitle="by MakerForce">
      <template v-slot:buttons-right>
        <button v-if="running" class="my-2 ml-4" @click="toggleRunning">
          <fa-icon size="2x" :icon="['far', 'pause']" />
        </button>
        <button v-if="completed" class="my-2 ml-4" @click="toggleRunning">
        </button>
        <button v-else class="my-2 ml-4" @click="toggleRunning">
          <fa-icon size="2x" :icon="['far', 'play']" />
        </button>
      </template>
      <template v-slot:content>
        <div class="mx-3 my-6">
          <h2 class="font-medium">Loss</h2>
          <trend
            :data="major"