<script setup lang="ts">
import AccretionIcon from "@/components/AccretionIcon/AccretionIcon.vue";
</script>

<template>
    <div class="container">
        <input type="text" placeholder="Search for accreted resources" />
        <div class="search-icon">
            <button class="search-icon__button" data-test-id="search-icon">
                <AccretionIcon kind="search"></AccretionIcon>
            </button>
        </div>
        <div class="clear-icon">
            <button class="clear-icon__button" data-test-id="close-icon">
                <AccretionIcon kind="close"></AccretionIcon>
            </button>
        </div>
    </div>
</template>

<style lang="scss" scoped>
@use "sass:math";
@import "../../theme/size";
@import "../../theme/colour";

.container {
    position: relative;

    input {
        background-color: darken($light-background, 2);
        padding: math.div($grid, 2) $grid * 2;
        width: $grid * 20;
        color: $light-body-font;
        font-size: 1.5rem;
        border: 1px solid darken($light-background, 4);

        @media (prefers-color-scheme: dark) {
            background-color: darken($dark-background, 2);
            color: $dark-body-font;
            border: 1px solid darken($dark-background, 4);
        }

        &:focus {
            outline: none;
            border-color: $light-body-font;

            @media (prefers-color-scheme: dark) {
                border-color: $dark-body-font;
            }
        }

        &::placeholder {
            color: rgba($light-body-font, 0.5);

            @media (prefers-color-scheme: dark) {
                color: rgba($dark-body-font, 0.5);
            }
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
        color: $light-body-font;
        cursor: pointer;

        @media (prefers-color-scheme: dark) {
            color: $dark-body-font;
        }
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
        color: $light-body-font;
        cursor: pointer;

        @media (prefers-color-scheme: dark) {
            color: $dark-body-font;
        }
    }
}
</style>
