export interface MediaServerParams {
  pageNo: number;
  pageSize: number;
}

export interface MediaServerUpdateParams {
  name: string;
  serverIp: string;
  serverPort: string;
  remark: string;
}
