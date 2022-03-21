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
          87%
        </CenteredText>
      </Card>
    </Cards>
  </div>
</template>

<script>
import Header from '~/components/common/Header.vue'
import Cards from '~/components/common/Cards.vue'
import Card from '~/components/common/Card.vue'
import CenteredText from '~/components/common/CenteredText.vue'
import trend from 'vuetrend'

const gradient = ['#ffffff', '#ff974d']

export default {
  layout: 'client',
  components: {
    Header,
    Cards,
    Card,
    CenteredText,
    trend
  },
  computed: {
    shareLink() {
      return 'https://staging.webml.app/' + this.$route.params.id
    }
  },
  data() {
    return {
      gradient,
      title: 'MNIST',
      major: [0, 0, 0, 0]
    }
  },
  created: function() {
    const base = process.env.NUXT_ENV_BACKEND2_URL || 'http://localhost:10200'
    const id = this.$route.params.id
    fetch(`${base}/metadata?model=${id}`)
      .then(() => {
        return res.json()
      })
      .then(body => {
        this.title = body.title
      })
  }
}
</script>
