<script setup lang="ts"></script>

<template>
    <div class="container">
        <slot />
    </div>
</template>

<style lang="scss" scoped>
@import "../../theme/size";

.container {
    padding: $grid;
}
</style>
