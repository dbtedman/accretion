<script setup lang="ts">
import { RouterView } from "vue-router";
import AccretionSidebar from "./components/AccretionSidebar/AccretionSidebar.vue";
import AccretionLayout from "@/components/AccretionLayout/AccretionLayout.vue";
import AccretionLayoutContent from "@/components/AccretionLayoutContent/AccretionLayoutContent.vue";
</script>

<template>
    <AccretionLayout>
        <AccretionSidebar />
        <AccretionLayoutContent>
            <RouterView />
        </AccretionLayoutContent>
    </AccretionLayout>
</template>
