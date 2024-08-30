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
  <div class="flex flex-col items-center justify-center h-full w-full">
    <p class="text-2xl font-bold my-8 text-black">Sube una receta médica</p>
    <FileInput @fileUploaded="handleFileUploaded" />
    <div v-if="result" class="mt-4 text-lg flex flex-col justify-center items-center">
      <PulseLoader v-if="loading" color="#2563EB" class="my-4" />
      <p class="p-8 my-8 w-full text-black rounded border-2 border-gray-300" v-html="result" v-else></p>
      <img v-if="imageUrl" :src="imageUrl" alt="Uploaded image" class="max-h-96" />
    </div>
  </div>
</template>

<style scoped>
/* Estilos opcionales */
</style>

