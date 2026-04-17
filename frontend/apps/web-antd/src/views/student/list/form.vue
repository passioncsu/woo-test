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
        allowClear: true,
        options: [
          { label: '男', value: '男' },
          { label: '女', value: '女' },
        ],
      },
      fieldName: 'gender',
      label: '性别',
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
      fieldName: 'className',
      label: '班级',
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
      componentProps: { placeholder: '请输入地址' },
      fieldName: 'address',
      label: '地址',
    },
  ],
  commonConfig: {
    componentProps: {
      class: 'w-full',
    },
  },
  layout: 'horizontal',
  wrapperClass: 'grid-cols-1 md:grid-cols-2',
});

async function handleSubmit() {
  const result = await formApi.validate();
  if (Object.keys(result?.errors ?? {}).length > 0) {
    return;
  }

  const rawValues = formApi.getValues();

  // 清理空值，避免发送无意义的空字符串
  const values: Record<string, any> = {};
  for (const [key, value] of Object.entries(rawValues)) {
    if (value !== undefined && value !== null && value !== '') {
      values[key] = value;
    }
  }

  // 字段映射: className -> class（避免JS保留字冲突）
  if (values.className !== undefined) {
    values['class'] = values.className;
    delete values.className;
  }

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
  // 字段映射: class -> className（避免JS保留字冲突）
  const mapped = { ...data };
  if (mapped['class'] !== undefined) {
    mapped.className = mapped['class'];
    delete mapped['class'];
  }
  formApi.setValues(mapped);
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
