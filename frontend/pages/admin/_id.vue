<template>
  <div>
    <Header :title="title" subtitle="by MakerForce">
      <template v-slot:buttons-left>
        <router-link class="my-2 mr-4" :to="{ name: '