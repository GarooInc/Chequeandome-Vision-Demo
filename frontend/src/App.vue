<script setup>
import { ref } from 'vue';
import FileInput from './components/FileInput.vue';
import { analyzeImage } from './services/openai';
import PulseLoader from 'vue-spinner/src/PulseLoader.vue';

const imageUrl = ref(null);
const result = ref('');
const loading = ref(false);

const handleFileUploaded = async (url) => {
  imageUrl.value = url;
  result.value = 'Analizando...';
  loading.value = true;

  try {
    const response = await analyzeImage(url);
    result.value = response;
    loading.value = false;
  } catch (error) {
    result.value = 'Error al analizar la imagen.';
    console.error(error);
    loading.value = false;
  }
};
</script>

<template>
  <div class="flex flex-col items-center justify-around h-full w-full">
    <img src="/garoo.png" alt="Garoo logo" class="logo" />
    <p class="text-2xl font-bold my-4 text-black">Sube una receta médica</p>
    <div v-if="result" class="my-4 text-lg flex flex-col justify-center items-center">
      <PulseLoader v-if="loading" color="#2563EB" class="my-4" />
      <p class="p-8 my-8 w-full text-black rounded border-2 border-gray-300" v-html="result" v-else></p>
      <img v-if="imageUrl" :src="imageUrl" alt="Uploaded image" class="max-h-64" />
    </div>
    <FileInput @fileUploaded="handleFileUploaded" />
  </div>
</template>

<style scoped>
.logo {
  width: clamp(200px, 40%, 300px);
}
</style>

