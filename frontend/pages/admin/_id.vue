<template>
  <div>
    <Header :title="title" subtitle="by MakerForce">
      <template v-slot:buttons-left>
        <router-link class="my-2 mr-4" :to="{ name: 'admin' }">
          <fa-icon size="2x" :icon="['far', 'arrow-left']" />
        </router-link>
      </template>
      <template v-slot:buttons-right>
        <button class="my-2 ml-4">
          <fa-icon size="2x" :icon="['far', 'times-hexagon']" />
        </button>
        <button class="my-2 ml-4">
          <fa-icon size="2x" :icon="['far', 'pause']" />
        </button>
      </template>
      <template v-slot:content>
        <div class="mx-3 my-6">
          <h2 cla