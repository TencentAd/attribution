package com.attribution.datacube.common.flatten.record;

import lombok.AllArgsConstructor;

import java.io.Serializable;

@AllArgsConstructor
public abstract class FlattenedRecord implements Serializable {
    private static final long serialVersionUID = 1111014L;

    public abstract long getEventTime();
}
