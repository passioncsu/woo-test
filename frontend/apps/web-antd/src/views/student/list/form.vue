<script lang="ts" setup>
import { message } from 'ant-design-vue';
import { ref } from 'vue';

import { useVbenForm } from '#/adapter/form';
import { createStudentApi, updateStudentApi } from '#/api/modules/student';

const emit = defineEmits<{
  success: [];
}>();

const isEdit = ref(false);
const editId = ref<number>();

const [Form, formApi] = useVbenForm({
  schema: [
    {
      component: 'Input',
      componentProps: { placeholder: '请输入学号' },
      fieldName: 'student_no',
      label: '学号',
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: { placeholder: '请输入姓名' },
      fieldName: 'name',
      label: '姓名',
      rules: 'required',
    },
    {
      component: 'Select',
      componentProps: {
        placeholder: '请选择性别',
        options: [
          { label: '男', value: '男' },
          { label: '女', value: '女' },
        ],
      },
      fieldName: 'gender',
      label: '性别',
    },
    {
      component: 'DatePicker',
      componentProps: { placeholder: '请选择出生日期', class: 'w-full' },
      fieldName: 'birthday',
      label: '出生日期',
    },
    {
      component: 'Input',
      componentProps: { placeholder: '请输入手机号' },
      fieldName: 'phone',
      label: '手机号',
    },
    {
      component: 'Input',
      componentProps: { placeholder: '请输入邮箱' },
      fieldName: 'email',
      label: '邮箱',
    },
    {
      component: 'Input',
      componentProps: { placeholder: '请输入专业' },
      fieldName: 'major',
      label: '专业',
    },
    {
      component: 'Input',
      componentProps: { placeholder: '请输入班级' },
      fieldName: 'class',
      label: '班级',
    },
    {
      component: 'Input',
      componentProps: { placeholder: '请输入地址' },
      fieldName: 'address',
      label: '地址',
    },
    {
      component: 'Select',
      componentProps: {
        placeholder: '请选择状态',
        options: [
          { label: '在读', value: 1 },
          { label: '休学', value: 2 },
          { label: '毕业', value: 3 },
        ],
      },
      fieldName: 'status',
      label: '状态',
      defaultValue: 1,
    },
  ],
  commonConfig: {
    componentProps: {
      class: 'w-full',
    },
  },
});

async function handleSubmit() {
  const result = await formApi.validate();
  if (Object.keys(result?.errors ?? {}).length > 0) {
    return;
  }

  const values = formApi.getValues();

  if (isEdit.value && editId.value) {
    await updateStudentApi(editId.value, values);
    message.success('更新成功');
  } else {
    await createStudentApi(values as any);
    message.success('创建成功');
  }

  emit('success');
}

function setFormData(data: Record<string, any>) {
  isEdit.value = !!data.id;
  editId.value = data.id;
  formApi.setValues(data);
}

function resetFields() {
  isEdit.value = false;
  editId.value = undefined;
  formApi.resetForm();
}

defineExpose({ setFormData, resetFields });
</script>

<template>
  <div class="p-4">
    <Form />
    <div class="mt-4 flex justify-end gap-3">
      <a-button @click="resetFields">重置</a-button>
      <a-button type="primary" @click="handleSubmit">提交</a-button>
    </div>
  </div>
</template>
