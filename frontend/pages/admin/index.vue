
<template>
  <div>
    <Header title="Sessions">
      <template v-slot:buttons-right>
        <button class="my-2 ml-4" @click="showNewSessionDialog = true">
          <fa-icon size="2x" :icon="['far', 'plus']" />
        </button>
      </template>
    </Header>
    <Cards>
      <Card
        v-for="model in models"
        :title="model.title"
        :key="model.title"
        arrow
        :to="{ name: 'admin-id', params: { id: model.id } }"
      >
        <div class="flex">
          <Subcard subtitle="Elapsed">
            <CenteredText class="text-4xl">
              {{ model.elapsed }}
            </CenteredText>
          </Subcard>
          <Subcard subtitle="Loss">
            <CenteredText class="text-4xl">
              {{ model.loss }}
            </CenteredText>
          </Subcard>
        </div>
      </Card>
    </Cards>
    <NewSessionDialog :show.sync="showNewSessionDialog" @needsRefresh="fetch" />
  </div>
</template>

<script>
import Header from '~/components/common/Header.vue'
import Cards from '~/components/common/Cards.vue'