<template>
  <div>
    <Header :title="title" subtitle="by MakerForce">
      <template v-slot:buttons-left>
        <router-link class="my-2 mr-4" :to="{ name: 'admin' }">
          <fa-icon size="2x" :icon="['far', 'arrow-left']" />
        </router-link>
      </template>
      <template v-slot:buttons-ri