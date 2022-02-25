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
          <h2 class="font-medium">Loss</h2>
          <trend
            :data="major"
            :gradient="gradient"
            :height="200"
			:width="400"
            auto-draw
            smooth
          >
          </trend>
        </div>
        <div class="mx-4 my-8">
          <div class="font-medium">Share this link</div>
          <div class="font-bold text-xl">{{ shareLink }}</div>
        </div>
      </template>
    </Header>
    <Cards>
      <Card subtitle="Elapsed time">
        <CenteredText class="text-4xl">
		Completed
        </CenteredText>
      </Card>
      <Card subtitle="Loss">
        <CenteredText class="text-4xl">
          0.567
        </CenteredText>
      </Card>
      <Card subtitle="Batch No.">
        <CenteredText class="text-4xl">
          468
        </CenteredText>
      </Card>
      <Card subtitle="Accuracy">
        <CenteredText class="text-4xl">