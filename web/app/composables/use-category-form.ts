import { defaultCategory } from "~/schemas/category-schema";

export const useCategoryForm = () => {
  const state = reactive(defaultCategory);

  return { state };
};
