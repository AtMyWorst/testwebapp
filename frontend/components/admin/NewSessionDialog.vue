
<template>
  <div
    class="fader fixed inset-0 p-6 flex items-center justify-center"
    v-if="show"
  >
    <div class="bg-white rounded-card shadow-md m-6">
      <div
        class="bg-primary text-white rounded-card flex items-center justify-start"
      >
        <button class="m-5" @click="$emit('update:show', false)">
          <fa-icon size="2x" :icon="['far', 'times']" />
        </button>
        <h1 class="text-2xl font-bold flex-1">New Session</h1>
        <button class="m-5 p-1" @click="submit">
          Create
        </button>
      </div>
      <div class="p-2">
        <div class="m-4">
          <input
            class="text-2xl font-bold"
            v-model="sessionTitle"
            placeholder="Enter session name"
          />
        </div>
        <div class="m-4">
          <DropArea
            @input="sessionModel = $event"
            placeholder="Drop model here"
            label="Model"
          />
        </div>
        <div class="m-4">
          <DropArea
            @input="sessionData = $event"
            placeholder="Drop data here"
            label="Data"
          />
        </div>
        <div class="m-4">
          <DropArea
            @input="sessionLabels = $event"
            placeholder="Drop labels here"
            label="Labels"
          />
        </div>
        <div class="m-4">
          <DropArea
            @input="sessionDataParser = $event"
            placeholder="Drop data parser here"
            label="Data Parser"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.fader {
  background: rgba(0, 0, 0, 0.5);
}
</style>

<script>
import DropArea from '~/components/common/DropArea.vue'
import cuid from 'cuid'
export default {
  components: {
    DropArea
  },
  props: {
    show: Boolean
  },
  data() {
    return {
      sessionTitle: '',
      sessionModel: null,
      sessionData: null,
      sessionLabels: null,
      sessionDataParser: null
    }
  },
  methods: {
    async submit() {
      // let files = this.$refs.fileloader.files

      // validate
      if (
        !this.sessionTitle ||
        !this.sessionModel ||
        !this.sessionData ||
        !this.sessionLabels ||
        !this.sessionDataParser
      ) {
        return
      }

      // TODO: create meta.json

      const base = process.env.NUXT_ENV_BACKEND2_URL || 'http://localhost:10200'
      const modelid = cuid()
      const dataid = cuid()
      const dataparserid = cuid()
      await fetch(base + '/model?id=' + modelid, {
        method: 'PUT',
        redirect: 'follow',
        body: this.sessionModel
      })
      await fetch(base + '/data?model=' + modelid + '&id=' + dataid, {
        method: 'PUT',
        redirect: 'follow',
        body: this.sessionData
      })
      await fetch(base + '/labels?model=' + modelid + '&id=' + dataid, {
        method: 'PUT',
        redirect: 'follow',
        body: this.sessionLabels
      })
      await fetch(base + '/data_parser?model=' + modelid + '&id=' + modelid, {
        method: 'PUT',
        redirect: 'follow',
        body: this.sessionDataParser
      })
      await fetch(base + '/metadata?model=' + modelid, {
        method: 'PUT',
        redirect: 'follow',
        body: JSON.stringify({
          title: this.sessionTitle
        })
      })
      await fetch(
        base +
          '/batch?model_id=' +
          modelid +
          '&data_id=' +
          dataid +
          '&batch_size=500',
        {
          method: 'POST',
          redirect: 'follow'
        }
      )
      this.$emit('update:show', false)
      this.$emit('needRefresh')
    }
  }
}
</script>