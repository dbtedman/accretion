<script setup lang="ts"></script>

<template>
    <div class="container"></div>
</template>

<style lang="scss" scoped>
@use "sass:math";
@import "../../theme/colour";
@import "../../theme/size";

$sidebar-dark-background: lighten($dark-background, 2);
$sidebar-light-background: darken($light-background, 2);
$sidebar-dark-border: lighten($dark-background, 4);
$sidebar-light-border: darken($light-background, 4);
$size-shadow: math.div($grid, 4);

.container {
    background-color: $sidebar-light-background;
    box-shadow: 0 0 $size-shadow rgba($sidebar-light-background, 0.5);
    padding: $grid;
    border-right: 1px solid $sidebar-light-border;
    width: $grid * 12;

    @media (prefers-color-scheme: dark) {
        background-color: $sidebar-dark-background;
        box-shadow: 0 0 $size-shadow rgba($sidebar-dark-background, 0.5);
        border-right-color: $sidebar-dark-border;
    }
}
</style>
