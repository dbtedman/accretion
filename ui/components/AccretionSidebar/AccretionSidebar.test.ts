import { expect, test } from "vitest";
import { mount } from "@vue/test-utils";
import AccretionSidebar from "@/components/AccretionSidebar/AccretionSidebar.vue";

test("renders default slot inside container", () => {
    const wrapper = mount(AccretionSidebar, {});
    expect(wrapper.find(".container").text()).toEqual("");
});
