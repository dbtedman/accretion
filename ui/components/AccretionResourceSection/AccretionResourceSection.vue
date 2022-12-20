<script setup lang="ts"></script>

<template>
    <section class="container">
        <h1>Section</h1>
        <div>
            <slot />
        </div>
    </section>
</template>

<style lang="scss" scoped>
@use "sass:math";
@import "../../theme/size";
@import "../../theme/colour";

.container {
    padding: math.div($grid, 2) $grid;
    margin-top: $grid;
    margin-bottom: $grid;
    border: 1px solid darken($light-background, 4);
    @media (prefers-color-scheme: dark) {
        border: 1px solid darken($dark-background, 4);
    }

    h1 {
        margin: 0;
        padding: 0;
        font-size: 1.2rem;
        font-weight: 600;
    }
}
</style>
