<script setup>

import FileInput from "./components/FileInput.vue";

import { ref } from 'vue';

const handleFileUploaded = (url) => {
  console.log("File uploaded", url);
}

</script>

<template>

  <div class="flex flex-col items-center justify-center h-screen">
    <p class="text-2xl font-bold my-2">Upload a file</p>
    <FileInput @fileUploaded="handleFileUploaded" />
  </div>

</template>

<style scoped>

</style>
