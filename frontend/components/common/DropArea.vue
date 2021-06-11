
<template>
  <label
    class="block text-center font-medium p-6"
    :class="{
      'bg-gray-100': !!filename,
      'text-gray-800': !!filename,
      'text-gray-500': !filename,
      'bg-gray-200': !filename
    }"
    @drop="dropInput"
    @dragover="dragOver"
    :for="id"
  >
    {{ filename ? label + ': ' + filename : placeholder }}
    <input class="hidden" @input="clickInput" type="file" :id="id" ref="file" />
  </label>
</template>

<script>
export default {
  props: {
    label: String,
    placeholder: String
  },
  data() {
    return {
      id: Math.floor(Math.random() * 10000000),
      filename: ''
    }
  },
  methods: {
    dropInput(e) {
      e.preventDefault()
      try {
        const item = e.dataTransfer.items[0]
        if (!item) {
          throw new Error('Item not found')
        }
        const file = item.getAsFile()
        this.chosenOne(file)
      } catch (e) {
        this.filename = ''
      }
    },
    dragOver(e) {
      e.preventDefault()
    },
    clickInput(e) {
      try {
        const file = e.target.files[0]
        if (!file) {
          throw new Error('Item not found')
        }
        this.chosenOne(file)
      } catch (e) {
        this.filename = ''
      }
    },
    chosenOne(file) {
      this.filename = file.name
      console.log(file)
      this.$emit('input', file)
    }
  }
}
</script>