package main;

import org.jooq.DataType;
import org.jooq.Field;
import org.jooq.TableField;
import org.jooq.impl.SQLDataType;

import java.util.Dictionary;
import java.util.Map;

public class Attribute {
    private final String name;
    private DataType<?> type;
    private final String table;
    private Attribute src = null;


    public Attribute(String name, DataType<?> type, String table) {
        this.name = name;
        this.type = type;
        this.table = table;
    }

    public Attribute(String name, DataType<?> type, String table, Attribute src) {
        this.name = name;
        this.type = type;
        this.table = table;
        this.src = src;
    }

    public void setType(DataType<?> type) {
        this.type = SQLDataType.VARCHAR;
    }

    public String getName() {
        return name;
    }

    public DataType<?> getType(){
        return type;
    }

    public String getTable() {
        return table;
    }

    public Attribute getSrc() {
        return src;
    }
}
