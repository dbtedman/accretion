import { expect, test } from "vitest";
import { mount } from "@vue/test-utils";
import AccretionLayout from "@/components/AccretionLayout/AccretionLayout.vue";

test("renders default slot inside container", () => {
    const wrapper = mount(AccretionLayout, {
        slots: {
            default: "main content",
        },
    });
    expect(wrapper.find(".container").html()).toContain("main content");
});
