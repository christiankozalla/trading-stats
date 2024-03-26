<template>
  <Panel :pt="{ header: { class: 'datapanel-header' }, content: { class: 'datapanel-content' } }">
    <slot></slot>
    <slot name="left"></slot>
    <slot name="right"></slot>
  </Panel>
</template>

<style>
.datapanel-header {
  padding: 9px 18px;
}

.datapanel-content > * {
  display: inline-block;
  width: 50%;
  vertical-align: top;
}

.datapanel-content > *:nth-child(1) {
  padding-right: 9px;
}

.datapanel-content > *:nth-child(2) {
  padding-left: 9px;
}
.datapanel-content > *:only-child {
  width: 100%;
  padding: 0;
}
</style>
