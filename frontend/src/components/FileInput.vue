<script setup>
import { ref } from 'vue';

import {PaperClipIcon} from "@heroicons/vue/24/solid";

const emit = defineEmits(['fileUploaded']);

const fileInput = ref(null);

function handleFileUpload() {
  const file = fileInput.value.files[0];
  if (file) {
    const formData = new FormData();
    formData.append('files', file);

    fetch('https://chequeandome-garoo.koyeb.app/files', {
      method: 'POST',
      body: formData,
    })
        .then(response => {
          const r = response.json()
          console.log(r)
          return r
        })
        .then(data => {
          console.log("Data", data)
          if (data.code === 200) {
            // Emit the URL back to the parent component
            console.log("File uploaded", data.paths[0])
            emit('fileUploaded', data.paths[0]);
          } else {
            console.error('File upload failed:', data.error);
          }
        })
        .catch(error => {
          console.error('Error:', error);
        });
  }
}

</script>

<template>
  <div class="w-1/2">
    <input type="file" ref="fileInput" style="display: none" @change="handleFileUpload" />
    <button @click="fileInput.click()" class="flex justify-center items-center text-white font-bold py-2 px-4 rounded w-full bg-blue-500  hover:bg-blue-700">
      <PaperClipIcon class="w-7 h-7" />
    </button>
  </div>
</template>

<style scoped>

</style>