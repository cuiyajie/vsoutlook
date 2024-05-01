export enum Signal {
  OpenNewTemplate,
  OpenNewTmplType,
  OpenTemplateImport,
  OpenNewSystem,
  OpenConfirmDialog,
  OpenResourceConfig,
  OpenAlertDialog,
  OpenNewUser,
  OpenNodeEdit,
  OpenTmplMetaEdit,
  OpenDownloadConfig,
  OpenNewLayout,
  OpenLayoutCellSetting,
  LayoutSaveAs,
}

export enum DeviceStatus {
  Normal = 1,
  Shutdown
}

export enum InputSignal {
  Normal = 1,
  None
}

export enum OutputSignal {
  Normal = 1,
  None
}

export enum PtpStatus {
  Locked = 1,
  Unlocked
}

export enum MirrorType {
  Container = 1,
  Virtual
}

export enum UserRole {
  Admin = 1,
  Normal = 2
}

export enum LayoutSize {
  FK = 1,
  HD = 2
}
