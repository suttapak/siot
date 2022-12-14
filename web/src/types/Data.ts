export interface DataControl {
  id: number;
  createdAt: Date;
  updatedAt: Date;
  data: number;
  label: string;
  controlId: number;
}

export interface DataDisplay {
  id: number;
  createdAt: Date;
  updatedAt: Date;
  data: number;
  label: string;
  displayId: number;
}
