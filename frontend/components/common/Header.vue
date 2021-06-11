
<template>
  <div class="head bg-primary text-white rounded-b-card p-2">
    <div class="max-w-4xl m-auto">
      <div class="flex items-start p-4">
        <slot name="buttons-left"> </slot>
        <div class="flex-grow">
          <h1 class="font-bold text-4xl">
            {{ title }}
          </h1>
          <div class="font-bold">
            {{ subtitle }}
          </div>
        </div>
        <slot name="buttons-right"> </slot>
      </div>