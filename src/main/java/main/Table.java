package main;

import org.jooq.*;
import org.jooq.Package;
import org.jooq.Record;

import java.sql.Connection;
import java.util.ArrayList;
import java.util.Collection;
import java.util.List;
import java.util.function.BiFunction;
import java.util.function.Function;
import java.util.stream.Stream;

public class Table {
    private final String name;
    private final ArrayList<Attribute> attributes;
    private final ArrayList<String> primaryKeys;
    private final ArrayList<String> foreignKeys;

    public Table(String name, ArrayList<Attribute> attributes, ArrayList<String> primaryKeys,
                 ArrayList<String> foreignKeys) {
        this.name = name;
        this.attributes = attributes;
        this.primaryKeys = primaryKeys;
        this.foreignKeys = foreignKeys;
    }

    public String getName() {
        return name;
    }

    public ArrayList<Attribute> getAttributes() {
        return attributes;
    }

    public ArrayList<String> getPrimaryKeys() {
        return primaryKeys;
    }

    public ArrayList<String> getForeignKeys() {
        return foreignKeys;
    }

    public void createTable(DSLContext dsl) {
        System.out.println(dsl.createTable(name).getSQL());
        CreateTableElementListStep tab = dsl.createTable(name);
        for (Attribute attribute : attributes) {
//            System.out.println(dsl.alterTable(name).addColumn(attribute.getName(), attribute.getType()).getSQL());
            tab = tab.column(attribute.getName(), attribute.getType());
        }
        if (primaryKeys != null && !primaryKeys.isEmpty()) {
            tab = tab.primaryKey(primaryKeys.toArray(new String[0]));
        }
        tab.execute();
    }
}
