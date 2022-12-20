import { expect, test } from "vitest";
import { mount } from "@vue/test-utils";
import AccretionLayoutContent from "@/components/AccretionLayoutContent/AccretionLayoutContent.vue";

test("renders default slot inside container", () => {
    const wrapper = mount(AccretionLayoutContent, {
        slots: {
            default: "main content",
        },
    });
    expect(wrapper.find(".container").html()).toContain("main content");
});
