<script setup lang="ts">
import { RouterView } from "vue-router";
import AccretionLayout from "@/components/AccretionLayout/AccretionLayout.vue";
import AccretionLayoutContent from "@/components/AccretionLayoutContent/AccretionLayoutContent.vue";
import AccretionSidebar from "@/components/AccretionSidebar/AccretionSidebar.vue";
</script>

<template>
    <AccretionLayout>
        <AccretionSidebar />
        <AccretionLayoutContent>
            <RouterView />
        </AccretionLayoutContent>
    </AccretionLayout>
</template>
