<script setup lang="ts"></script>

<template>
    <div class="container">
        <input type="text" placeholder="Search for accreted resources" />
        <div class="search-icon">
            <!-- https://fonts.google.com/icons -->
            <!-- TODO: Encapsulate icons in a component. -->
            <button class="search-icon__button">
                <span class="material-symbols-outlined">search</span>
            </button>
        </div>
        <div class="clear-icon">
            <!-- https://fonts.google.com/icons -->
            <!-- TODO: Encapsulate icons in a component. -->
            <button class="clear-icon__button">
                <span class="material-symbols-outlined">close</span>
            </button>
        </div>
    </div>
</template>

<style lang="scss" scoped>
@use "sass:math";
@import "../../theme/size";
@import "../../theme/colour";

// TODO: Dark/Light

.container {
    position: relative;

    input {
        background-color: darken($dark-background, 2);
        padding: math.div($grid, 2) $grid * 2;
        width: $grid * 20;
        color: $dark-body-font;
        font-size: 1.5rem;
        border: 1px solid darken($dark-background, 4);

        &:focus {
            outline: none;
            border-color: $dark-body-font;
        }

        &::placeholder {
            color: rgba($dark-body-font, 0.5);
        }
    }
}

.search-icon {
    position: absolute;
    left: 0;
    top: 0;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;

    &__button {
        padding: 0;
        margin-left: math.div($grid, 1.5);
        margin-right: math.div($grid, 1.5);
        background: none;
        border: none;
        color: $dark-body-font;
        cursor: pointer;
    }
}

.clear-icon {
    position: absolute;
    right: 0;
    top: 0;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;

    &__button {
        padding: 0;
        margin-left: math.div($grid, 1.5);
        margin-right: math.div($grid, 1.5);
        background: none;
        border: none;
        color: $dark-body-font;
        cursor: pointer;
    }
}
</style>
