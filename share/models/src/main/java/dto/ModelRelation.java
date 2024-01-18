package dto;

import dto.compositeKeys.ModelRelationKey;
import jakarta.persistence.*;
import lombok.AccessLevel;
import lombok.Builder;
import lombok.Getter;
import lombok.NoArgsConstructor;

import java.util.UUID;

@Entity
@Table(name = "model_relation")
@Getter
@NoArgsConstructor(access = AccessLevel.PROTECTED)
@IdClass(ModelRelationKey.class)
public class ModelRelation {
    @Id
    @Column(name = "model_id", nullable = false)
    private UUID modelId;
    @Id
    @Column(name = "child_model_id", nullable = false)
    private UUID childModelId;

    @ManyToOne
    @JoinColumn(name = "model_id", updatable = false, insertable = false)
    private Model model;

    @Builder
    public ModelRelation(UUID modelId, UUID childModelId) {
        this.modelId = modelId;
        this.childModelId = childModelId;
    }
}
