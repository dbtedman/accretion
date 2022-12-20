<script setup lang="ts">
// TODO: How to handle including the necessary external scripts.
// https://fonts.google.com/icons
export type AccretionIconType = "search" | "close";
const props = defineProps<{ kind: AccretionIconType }>();
</script>

<template>
    <span class="material-symbols-outlined">{{ props.kind }}</span>
</template>
