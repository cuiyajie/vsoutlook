export enum Signal {
  OpenNewTemplate,
  OpenNewTmplType,
  OpenTemplateImport,
  OpenNewSystem,
  OpenConfirmDialog,
  OpenResourceConfig
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
  Admin = 1
}
