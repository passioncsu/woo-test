<script lang="ts" setup>
import { Page, useVbenModal } from '@vben/common-ui';
import { message, Modal } from 'ant-design-vue';
import { nextTick, ref } from 'vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  deleteStudentApi,
  getStudentDetailApi,
  getStudentListApi,
} from '#/api/modules/student';

import StudentForm from './form.vue';

const formRef = ref();

const [FormModal, formModalApi] = useVbenModal({
  onOpenChange(isOpen) {
    if (isOpen) {
      formRef.value?.resetFields();
    }
  },
});

const statusMap: Record<number, string> = {
  1: '在读',
  2: '休学',
  3: '毕业',
};

const statusColorMap: Record<number, string> = {
  1: 'green',
  2: 'orange',
  3: 'blue',
};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions: {
    columns: [
      { type: 'seq', width: 60, title: '序号' },
      { field: 'student_no', title: '学号', width: 140 },
      { field: 'name', title: '姓名', width: 120 },
      { field: 'gender', title: '性别', width: 80 },
      { field: 'major', title: '专业', width: 160 },
      { field: 'class', title: '班级', width: 140, formatter: ({ cellValue }) => cellValue || '-' },
      { field: 'phone', title: '电话', width: 140 },
      { field: 'status', title: '状态', width: 100, slots: { default: 'status' } },
      { field: 'created_at', title: '创建时间', width: 180 },
      { field: 'action', title: '操作', width: 180, slots: { default: 'action' }, fixed: 'right' },
    ],
    proxyConfig: {
      ajax: {
        query: async ({ page }) => {
          const searchForm = gridApi.formApi.getValues();
          const res = await getStudentListApi({
            keyword: searchForm.keyword || '',
            page: page.currentPage,
            pageSize: page.pageSize,
          });
          return { items: res.list || [], total: res.total || 0 };
        },
      },
    },
    toolbarConfig: {
      search: true,
    },
  },
  formOptions: {
    schema: [
      {
        component: 'Input',
        componentProps: { placeholder: '请输入姓名/学号/专业' },
        fieldName: 'keyword',
        label: '关键词',
      },
    ],
  },
});

function handleAdd() {
  formModalApi.setState({ title: '新增学生' });
  formModalApi.open();
  nextTick(() => {
    formRef.value?.setFormData({});
  });
}

async function handleEdit(record: any) {
  formModalApi.setState({ title: '编辑学生' });
  formModalApi.open();
  const detail = await getStudentDetailApi(record.id);
  nextTick(() => {
    formRef.value?.setFormData(detail);
  });
}

function handleDelete(record: any) {
  Modal.confirm({
    title: '确认删除',
    content: `确定要删除学生「${record.name}」吗？`,
    async onOk() {
      await deleteStudentApi(record.id);
      message.success('删除成功');
      gridApi.reload();
    },
  });
}

function handleFormSuccess() {
  formModalApi.close();
  gridApi.reload();
}
</script>

<template>
  <Page auto-content-height>
    <Grid>
      <template #toolbar-tools>
        <a-button type="primary" @click="handleAdd">
          新增学生
        </a-button>
      </template>

      <template #status="{ row }">
        <a-tag :color="statusColorMap[row.status] || 'default'">
          {{ statusMap[row.status] || '未知' }}
        </a-tag>
      </template>

      <template #action="{ row }">
        <a-button type="link" size="small" @click="handleEdit(row)">编辑</a-button>
        <a-button type="link" danger size="small" @click="handleDelete(row)">删除</a-button>
      </template>
    </Grid>

    <FormModal>
      <StudentForm ref="formRef" @success="handleFormSuccess" />
    </FormModal>
  </Page>
</template>
