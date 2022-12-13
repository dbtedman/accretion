<script setup lang="ts"></script>

<template>
    <div class="container">
        <slot />
    </div>
</template>

<style lang="scss" scoped>
@import "../../theme/colour";
.container {
    width: 100vw;
    height: 100vh;
    display: flex;
}
</style>
