declare module '@apiverve/stateboundaries' {
  export interface stateboundariesOptions {
    api_key: string;
    secure?: boolean;
  }

  /**
   * Describes fields the current plan does not unlock. Locked fields arrive as null
   * in `data`; `locked_fields` names them, using dot paths for nested fields.
   * Absent when the plan unlocks everything.
   */
  export interface PremiumInfo {
    message: string;
    upgrade_url: string;
    locked_fields: string[];
  }

  export interface stateboundariesResponse {
    status: string;
    error: string | null;
    data: StateBoundariesData;
    code?: number;
    premium?: PremiumInfo;
  }


  interface StateBoundariesData {
      type:     null | string;
      features: Feature[];
  }
  
  interface Feature {
      type:       null | string;
      geometry:   Geometry;
      properties: Properties;
  }
  
  interface Geometry {
      coordinates: Array<Array<Array<(number | null)[]>>>;
      type:        null | string;
  }
  
  interface Properties {
      gid:        number | null;
      arealand:   null | string;
      division:   number | null;
      intptlat:   number | null;
      name:       null | string;
      objectid:   number | null;
      areawater:  null | string;
      intptlon:   number | null;
      oid:        null | string;
      funcstat:   null | string;
      centlon:    number | null;
      stusab:     null | string;
      state:      null | string;
      statens:    null | string;
      centlat:    number | null;
      basename:   null | string;
      mtfcc:      null | string;
      region:     number | null;
      lsadc:      null | string;
      geoid:      null | string;
      geoPoint2D: GeoPoint2D;
  }
  
  interface GeoPoint2D {
      lon: number | null;
      lat: number | null;
  }

  export default class stateboundariesWrapper {
    constructor(options: stateboundariesOptions);

    execute(callback: (error: any, data: stateboundariesResponse | null) => void): Promise<stateboundariesResponse>;
    execute(query: Record<string, any>, callback: (error: any, data: stateboundariesResponse | null) => void): Promise<stateboundariesResponse>;
    execute(query?: Record<string, any>): Promise<stateboundariesResponse>;
  }
}
